package redis

import (
	"app/models"
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math"
	"time"
)

/*
direction=1时：
	1，之前没有投过票，现在投赞成票   差值的绝对值1
	2，之前投反对票，现在改赞成票    差值的绝对值2
direction=0时：
	1，之前投赞成票，现在取消投票   差值的绝对值1
	2，之前投反对票，现在取消投票   差值的绝对值1
direction=-1时：
	1，之前没有投过票，现在投反对票   差值的绝对值1
	2，之前投过赞成票，现在改投反对票  差值的绝对值2

投票限制：
	每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许用户再投票
	超过一个星期之后，将redis中保存的赞成票和反对票的记录存储到mysql表中，并将redis中的记录删除

投票分数：投一票加432分
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600 //一周时间(s)
	scorePreVote     = 432           //每一票分数值
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// CreatePost  创建帖子记录
func CreatePost(postID uint64) error {
	err := rdb.ZAdd(ctx, KeyPostTimeZset, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Err()
	return err
}

// VoteForPost 为帖子投票
func VoteForPost(userID, postID string, value float64) error {
	//判断投票的限制
	//获取帖子的发布时间
	postTime := rdb.ZScore(ctx, KeyPostTimeZset, postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {

		//帖子发布时间超过一周
		zap.L().Error("the duration of the post exceeds one week", zap.Error(ErrVoteTimeExpire))
		return ErrVoteTimeExpire
	}

	//更新分数和记录用户给帖子的投票记录操作是一个事务
	txPipeLine := rdb.TxPipeline()

	//更新分数
	//查询当前用户给当前帖子的投票记录
	ov := txPipeLine.ZScore(ctx, KeyPostVotedZsetPrefix+postID, userID).Val()

	//如果已经投票了，就不允许再次投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值绝对值
	txPipeLine.ZIncrBy(ctx, KeyPostScoreZset, op*diff*scorePreVote, postID)
	//if err != nil {
	//	zap.L().Error("rdb.ZIncrBy(ctx, KeyPostScoreZset, op*diff*scorePreVote, postID) failed", zap.Error(err))
	//	return err
	//}
	//记录用户为该帖子投票记录
	if value == 0 {
		txPipeLine.ZRem(ctx, KeyPostVotedZsetPrefix+postID, userID)

	} else {
		txPipeLine.ZAdd(ctx, KeyPostVotedZsetPrefix+postID, &redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err := txPipeLine.Exec(ctx)
	return err
}

// GetVoteDataList 根据帖子的id获取该帖子的赞成票数和反对票数
func GetVoteDataList(postIDAndScoreList []redis.Z) ([]*models.PostVoteData, error) {
	pipeLine := rdb.Pipeline()
	for _, post := range postIDAndScoreList {
		postIDStr, ok := post.Member.(string)
		if !ok {
			zap.L().Warn("post.Member.(string) failed")
			continue
		}
		pipeLine.ZCount(ctx, KeyPostVotedZsetPrefix+postIDStr, "1", "1")
		pipeLine.ZCount(ctx, KeyPostVotedZsetPrefix+postIDStr, "-1", "-1")
	}
	cmders, err := pipeLine.Exec(ctx)

	if err != nil {
		zap.L().Error("GetVoteDataList() failed", zap.Error(err))
		return nil, err
	}
	postVoteDataList := make([]*models.PostVoteData, len(postIDAndScoreList))
	for index := range postVoteDataList {
		postVoteDataList[index] = new(models.PostVoteData)
	}

	for index, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		if index%2 == 0 {
			//赞成票数
			postVoteDataList[index/2].SupportVoteNum = v
		} else {
			//反对票数
			postVoteDataList[index/2].UnsupportedVoteNum = v
		}
	}
	return postVoteDataList, nil
}
