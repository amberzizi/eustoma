package models

import (
	"errors"
)

//定义自定义error
var (
	ErrorVoteOutOfTime = errors.New("超出可投票时间")
	ErrorVoteRepeat    = errors.New("请勿重复投票")
)

//统一使用的常量名称
const CanVoteLimit = 3600 * 24 * 7 //允许投票的时间长度
const ScorePerVote = 432           //每一票价值的分数
const CheckForTime = 1             //按新旧帖子排序
const CheckForScore = 2            //按照帖子投票分数排序
//分类模型
//完整的分类信息
type Post struct {
	//Id             int64
	Post_id   int64 `json:"post_id,string"`
	Title     string
	Author_id int64 `json:"author_id,string"`
	//Introduction   string
	//Create_time    string
	//Update_time    string
}

type PostDetail struct {
	Id           int64 `json:"id,string"`
	Post_id      int64 `json:"post_id,string"`
	Title        string
	Content      string
	Author_id    int64 `json:"author_id,string"`
	Community_id int64 `json:"community_id,string"`
	Status       int
	Create_time  string
	Update_time  string
}

func (u Post) TableName() string {
	return "post"
}

func (u PostDetail) TableName() string {
	return "post"
}

//api返回帖子详情结构体
//嵌入社区分类信息结构体
//嵌入帖子结构体
type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*CommunityDetail `json:"community"`
	*PostDetail
}

//分数
type ApiPostDetailAndScore struct {
	Score           int64  `json:"score,string"`
	VoteNum         int64  `json:"votenum,string"`
	AuthorName      string `json:"author_name"`
	CommunityDetail `json:"community"`
	PostDetail
}
