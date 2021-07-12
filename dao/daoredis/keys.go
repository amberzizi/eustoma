package daoredis

//存储redis固定的key值

const (
	KeyPrefix              = "eustoma:"
	KeyPostTimeZSet        = "post:time"   //zset; 以发帖时间为分数
	KeyPostScoreZSet       = "post:score"  //zset；以投票为分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset;前缀   参数是帖子post_id 记录用户和投票类型
)
