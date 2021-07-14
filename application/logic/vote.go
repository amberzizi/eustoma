package logic

import (
	"errors"
	"mygin/application/models"
	"mygin/dao/daomysql"
	"time"
)

//定义自定义error
var (
	ErrorVoteOutOfTime = errors.New("超出可投票时间")
)

//简化版投票分数
//86400/200   需要200张赞成票 可以给帖子时间戳分数增加86400 挂在首页
/**
direction = 1：
（1）之前没有投票  赞成票
（2）之间投反对   现在
*/

//限制 可投票的时间限制 自发表之日起一个星期之内可以投票
//到期后将投票的zset 存储到mysql中 可查询 ； 删除zset
func PostVote(p *models.ParamVotePost, user_id int64) (bool, error) {
	//1.获取帖子信息
	postinfo, err := daomysql.GetPostDetailByPid(p.PostId)
	if err != nil {
		return false, err
	}
	//2.判断帖子时间是否可以投票  超出7天 报错超时
	stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", postinfo.Create_time, time.Local)
	createtime := stamp.Unix()
	nowtime := time.Now().Unix()
	if (createtime + models.CanVoteLimit) < nowtime {
		return false, ErrorVoteOutOfTime
	}
	return true, nil
}
