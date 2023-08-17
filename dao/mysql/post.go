package mysql

import (
	"app/models"
	"go.uber.org/zap"
)

// InsertPost 插入数据
func InsertPost(post *models.Post) (err error) {
	err = db.Omit("status", "create_time", "update_time").Create(post).Error
	if err != nil {
		zap.L().Debug("db.Create(post) failed", zap.Error(err))
		return
	}
	return
}

// GetPostDetailByPostID 根据帖子id查询数据
func GetPostDetailByPostID(postId int) (postDetail *models.Post, err error) {
	postDetail = new(models.Post)
	err = db.Where("post_id=?", postId).Find(postDetail).Error
	if err != nil {
		zap.L().Error("db.Where('post_id=?', postId).Find(postDetail) failed", zap.Error(err))
		return
	}
	return
}

// GetPostList 获取所有帖子信息
func GetPostList(page, pageSize int) (postList []*models.Post, err error) {
	postList = make([]*models.Post, 0, pageSize)
	err = db.Order("create_time desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&postList).Error
	return
}

// GetPostList2 根据分数获取帖子列表
func GetPostList2(page, pageSize int) (postList []*models.Post, err error) {
	postList = make([]*models.Post, 0, pageSize)
	err = db.Order("create_time desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&postList).Error
	return
}
