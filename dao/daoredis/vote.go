package daoredis

import (
	"github.com/go-redis/redis"
	"math"
	"mygin/application/models"
	"strings"
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
	//检查 用户重复投票，如果用户传递的投票值和之前相同，拒绝后续操作
	if value == ov {
		return false, models.ErrorVoteRepeat
	}

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

//获取符合要求的postid  最新/最高分
func GetPostListKeyvalueByParam(typeId int, cpage int, limit int) ([]string, []redis.Z, error) {
	//分页
	start_v := (cpage - 1) * limit
	stop_v := start_v + limit - 1
	//key索引
	key := getReidsKey(KeyPostTimeZSet) //models.CheckForTime
	if typeId == models.CheckForScore {
		key = getReidsKey(KeyPostScoreZSet)
	}
	//newest
	results, err := rdb.ZRevRange(key, int64(start_v), int64(stop_v)).Result()
	resultswithscore, err := rdb.ZRevRangeWithScores(key, int64(start_v), int64(stop_v)).Result()

	if err != nil {
		return nil, nil, err
	}
	return results, resultswithscore, nil
}

//获取指定帖子列表用户投的票数  赞成票 1
func GetPostListVoteNumByPostList(datas []string) (map[string]int64, error) {
	//returndata := make(map[string]int64)
	//for _, value := range datas {
	//	v := rdb.ZCount(getReidsKey(KeyPostVotedZSetPrefix+value), "1", "1").Val()
	//	returndata[value] = v
	//}

	//keys := make([]string, len(datas))
	//...
	//keys = append(keys, key)

	//使用pipeline 减少RTT连接时间
	pipeline := rdb.Pipeline()
	for _, id := range datas {
		key := getReidsKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")

	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	//创建返回map
	//根据pipeline合并执行的返回值 拼接返回map
	returndata := make(map[string]int64)
	for _, cmder := range cmders {
		postidstr := cmder.Args()[1].(string)
		desStr := strings.Split(postidstr, ":")
		v := cmder.(*redis.IntCmd).Val()
		returndata[desStr[3]] = v
	}
	return returndata, err

}
