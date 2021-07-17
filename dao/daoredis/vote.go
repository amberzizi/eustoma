package daoredis

import (
	"github.com/go-redis/redis"
	"math"
	"mygin/application/models"
	"time"
)

//投票
//redis关于投票有三哥zset的数据类型
//1.eustoma:post:voted:postId :记录了对于某帖子，谁投了什么票
//2.eustoma:post:score  :记录了帖子的分数
//3.eustoma:post:time   :记录了帖子的投票时间
//事务  3
func SaveVoteForPost(userId, postId string, value float64) (bool, error) {
	//先替换value
	if value == -99 {
		value = 0
	}

	//1.获取帖子发布时间  判断是否超出投票时间
	postTime := rdb.ZScore(getReidsKey(KeyPostTimeZSet), postId).Val()
	nowtime := time.Now().Unix()
	if float64(nowtime)-postTime > models.CanVoteLimit {
		return false, models.ErrorVoteOutOfTime
	}

	//2.更新帖子分数
	//查询之前的投票记录
	ov := rdb.ZScore(getReidsKey(KeyPostVotedZSetPrefix+postId), userId).Val()

	var op float64 //方向
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值  使用绝对值可以快读得到 加减分数的实际值

	//事务
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getReidsKey(KeyPostScoreZSet), op*diff*models.ScorePerVote, postId)
	//3.记录用户为该帖子投票的数据
	if value == 0 {
		//取消投票
		pipeline.ZRem(getReidsKey(KeyPostVotedZSetPrefix+postId), userId).Result()
	} else {
		pipeline.ZAdd(getReidsKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			value,
			userId,
		}).Result()
	}

	_, err := pipeline.Exec()

	return true, err

}

//保存发帖时的时间
//保存插入时间时间
//保存初始分数
//事务
func SavePostTimeAndInitScore(postId string) error {
	nottime := time.Now().Unix()
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getReidsKey(KeyPostTimeZSet), redis.Z{
		float64(nottime),
		postId,
	}).Result()
	pipeline.ZAdd(getReidsKey(KeyPostScoreZSet), redis.Z{
		float64(nottime),
		postId,
	}).Result()

	_, err := pipeline.Exec()

	return err
}
