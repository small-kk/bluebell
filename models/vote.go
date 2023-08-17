package models

// PostVoteData 帖子数据（赞成票、反对票）
type PostVoteData struct {
	SupportVoteNum     int64 `json:"support_vote_num"`
	UnsupportedVoteNum int64 `json:"unsupported_vote_num"`
}
