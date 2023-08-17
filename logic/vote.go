package logic

import (
	"app/dao/redis"
	"app/models"
	"strconv"
)

// VoteForPost 为帖子投票
func VoteForPost(userID int64, voteData *models.ParamVoteData) error {
	userIDStr := strconv.FormatInt(userID, 10)
	return redis.VoteForPost(userIDStr, voteData.PostID, float64(voteData.Direction))
}
