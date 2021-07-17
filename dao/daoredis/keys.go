package daoredis

//存储redis固定的key值

const (
	KeyPrefix              = "eustoma:"
	KeyPostTimeZSet        = "post:time"   //zset; 以发帖时间为分数  {value:postId,score:发帖时间}
	KeyPostScoreZSet       = "post:score"  //zset；以投票为分数  {value:postId,score:8000当前分数}
	KeyPostVotedZSetPrefix = "post:voted:" //zset;前缀   参数是帖子post_id  {value:userId,score:1 0 -1}
)

//获取带前缀的rediskey
func getReidsKey(key string) string {
	return KeyPrefix + key
}
