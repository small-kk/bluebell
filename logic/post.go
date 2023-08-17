package logic

import (
	"app/dao/mysql"
	"app/dao/redis"
	"app/models"
	"app/pkg/snowflake"
	"go.uber.org/zap"
	"strconv"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//生成postID
	p.PostID = uint64(snowflake.GenID())
	//保存在数据库
	err = mysql.InsertPost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.PostID)
	return
}

// GetPostDetailByPostID 根据贴子id查询帖子详情
func GetPostDetailByPostID(postId int) (apiPostDetail *models.ApiPostDetail, err error) {
	// 帖子id查询帖子内容
	postDetail, err := mysql.GetPostDetailByPostID(postId)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByPostID(postId) failed", zap.Error(err))
		return
	}
	//根据作者id查询作者详情
	user, err := mysql.GetUserByUserID(postDetail.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserNameByUserID(postDetail.AuthorID) failed", zap.Error(err))
		return
	}
	//根据社区分类id查询社区分类详情
	communityDetail, err := mysql.GetCommunityDetailByID(int(postDetail.CommunityID))
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(int(postDetail.CommunityID)) failed", zap.Error(err))
		return
	}

	apiPostDetail = &models.ApiPostDetail{
		AuthorName:         user.Username,
		Post:               postDetail,
		CommunityDetailRes: communityDetail,
	}
	return
}

// GetPostList 查询帖子列表
func GetPostList(page, pageSize int) (apiPostList []*models.ApiPostDetail, err error) {
	//获取帖子信息
	postList, err := mysql.GetPostList(page, pageSize)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return
	}
	apiPostList = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		//根据作者id获取作者信息
		user, err := mysql.GetUserByUserID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区分类id获取社区分类信息
		community, err := mysql.GetCommunityDetailByID(int(post.CommunityID))
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(int(post.CommunityID)) failed", zap.Error(err))
			continue
		}
		apiPost := &models.ApiPostDetail{
			AuthorName:         user.Username,
			Post:               post,
			CommunityDetailRes: community,
		}
		apiPostList = append(apiPostList, apiPost)
	}
	return
}

// GetPostList2 根据分数获取帖子列表
func GetPostList2(page, pageSize int) (apiPostList []*models.ApiPostDetail, err error) {
	//从redis中根据分数获取帖子id和分数列表
	postIDAndScoreList, err := redis.GetPostIDAndScoreList(page, pageSize)
	if err != nil {
		zap.L().Error("redis.GetPostIDList(page, pageSize) failed", zap.Error(err))
		return
	}

	//从redis中获取帖子赞成票数和反对票数
	VoteDataList, err := redis.GetVoteDataList(postIDAndScoreList)
	zap.L().Debug("VoteDataList, err := redis.GetVoteDataList(postIDAndScoreList)", zap.Any("VoteDataList", VoteDataList))
	if err != nil {
		zap.L().Error("redis.GetVoteDataList(postIDAndScoreList) failed", zap.Error(err))
		return
	}

	//apiPostList 包含用户名和社区分类的帖子详细信息
	apiPostList = make([]*models.ApiPostDetail, 0, len(postIDAndScoreList))

	for index, postIDAndScore := range postIDAndScoreList {
		//获取postIDStr
		postIDStr, ok := postIDAndScore.Member.(string)
		if !ok {
			zap.L().Warn("post.Member.(string) failed", zap.String("postID", postIDStr))
			return
		}
		//将postID转换成int
		var postID int
		postID, err = strconv.Atoi(postIDStr)
		if err != nil {
			zap.L().Error(" strconv.Atoi(postIDStr) failed", zap.Error(err))
			continue
		}
		//根据post_id获取post的详细信息
		var post *models.Post
		post, err = mysql.GetPostDetailByPostID(postID)
		if err != nil {
			zap.L().Error("mysql.GetPostDetailByPostID(postID) failed", zap.Error(err))
			continue
		}

		//根据作者id获取作者信息
		var user *models.User
		user, err = mysql.GetUserByUserID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区分类id获取社区分类信息
		var community *models.CommunityDetailRes
		community, err = mysql.GetCommunityDetailByID(int(post.CommunityID))
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(int(post.CommunityID)) failed", zap.Error(err))
			continue
		}
		//组装ApiPostDetail
		apiPost := &models.ApiPostDetail{
			AuthorName:         user.Username,
			PostTotalScore:     postIDAndScore.Score,
			PostVoteData:       VoteDataList[index],
			Post:               post,
			CommunityDetailRes: community,
		}
		apiPostList = append(apiPostList, apiPost)
	}
	return
}
