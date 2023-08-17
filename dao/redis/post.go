package redis

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// GetPostIDAndScoreList 根据帖子分数排序获取帖子id和分数列表
func GetPostIDAndScoreList(page, pageSize int) ([]redis.Z, error) {
	postListWithScores, err := rdb.ZRevRangeWithScores(ctx, KeyPostScoreZset, int64((page-1)*pageSize), int64(page*pageSize-1)).Result()
	zap.L().Debug("GetPostIDAndScoreList", zap.Any("postListWithScores", postListWithScores))
	return postListWithScores, err
}
