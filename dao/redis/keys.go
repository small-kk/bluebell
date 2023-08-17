package redis

// redis key

const (
	KeyPostTimeZset        = "bluebell:post:time"   //存储贴子的发帖时间
	KeyPostScoreZset       = "bluebell:post:score"  //存储帖子的分数
	KeyPostVotedZsetPrefix = "bluebell:post:voted:" //存储贴子的投票用户和用户投票类型（反对票或赞成票）
)
