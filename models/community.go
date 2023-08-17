package models

import "time"

// Community 结构体
type Community struct {
	CommunityID   uint64 `gorm:"column:community_id" json:"community_id"`
	CommunityName string `gorm:"column:community_name" json:"community_name"`
}

// CommunityDetailRes 社区详情
type CommunityDetailRes struct {
	CommunityID   uint64    `json:"community_id" gorm:"column:community_id"`
	CommunityName string    `json:"community_name" gorm:"column:community_name"`
	Introduction  string    `json:"introduce" gorm:"column:introduction"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}
