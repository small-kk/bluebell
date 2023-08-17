package models

import "time"

type Post struct {
	PostID      uint64    `json:"post_id" gorm:"column:post_id"`
	AuthorID    uint64    `json:"author_id" gorm:"column:author_id"`
	CommunityID uint64    `json:"community_id" gorm:"community_id" binding:"required"`
	Status      int32     `json:"status" gorm:"status"`
	Title       string    `json:"title" gorm:"title" binding:"required"`
	Content     string    `json:"content" gorm:"content" binding:"required"`
	CreateTime  time.Time `json:"-" gorm:"create_time"`
	UpdateTime  time.Time `json:"-" gorm:"update_time"`
}

// ApiPostDetail 帖子详情结构体
type ApiPostDetail struct {
	AuthorName          string                  `json:"author_name"`
	PostTotalScore      float64                 `json:"post_total_score"`
	*PostVoteData       `json:"post_vote_data"` //嵌入帖子数据
	*Post               `json:"post"`           //嵌入帖子结构体
	*CommunityDetailRes `json:"community"`      //嵌入社区分类结构体
}
