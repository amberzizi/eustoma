package logic

import (
	"mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/dao/daoredis"
	"mygin/tools/snowflake"
	"strconv"
)

//社区列表
func GetPostListByCid(cid int64, page int, limit int) ([]models.PostDetail, error) {
	datas, err := daomysql.GetPostListByCid(cid, page, limit)
	return datas, err
}

//获取最新或分数排序的帖子列表
func GetPostListIndexByParam(typeId int, cpage int, limit int) ([]models.ApiPostDetailAndScore, error) {
	//获取排序后的最新 或最高分postid数组
	datas, dataswitchscore, err := daoredis.GetPostListKeyvalueByParam(typeId, cpage, limit)
	//如果无相关帖子可显示 则不继续下面的查询
	if len(datas) == 0 {
		return nil, err
	}
	//获取本次帖子相关的redis正数
	dataforvotesum, err := daoredis.GetPostListVoteNumByPostList(datas)
	//fmt.Println(dataforvotesum)
	//到mysql中获取帖子详情
	datasdetaillist, err := daomysql.GetPostListPostIdList(datas)
	//初始化结构体数组
	var returninfo = make([]models.ApiPostDetailAndScore, len(dataswitchscore))

	//排序输出
	for key, valueZ := range dataswitchscore {
		for _, value := range datasdetaillist {
			if valueZ.Member.(string) == strconv.FormatInt(value.Post_id, 10) {
				//获取当前帖子的社区信息
				datacommunity, err := daomysql.GetCommunityByCid(value.Community_id)
				if err != nil {
					return nil, err
				}
				//获取当前帖子的用户信息
				datauser, err := daomysql.GetUserInfoByUserId(value.Author_id)
				if err != nil {
					return nil, err
				}
				//为当前帖子的正数票数赋值
				var votenumtemp int64
				for votsumpostid, votsumvalue := range dataforvotesum {
					if votsumpostid == valueZ.Member.(string) {
						votenumtemp = votsumvalue
					}
				}

				returninfo[key] = models.ApiPostDetailAndScore{
					AuthorName:      datauser.Username,
					VoteNum:         votenumtemp,         //帖子投票为正数赞成票的计数
					Score:           int64(valueZ.Score), //时间为检索的时候为时间戳，分数检索的时候代表分数
					PostDetail:      value,
					CommunityDetail: *datacommunity,
				}
				break
			}
		}
	}

	return returninfo, err
}

//社区内按时间和按分数排序
func GetPostListCommunityIndexByParam(communityid int64, typeId int, cpage int, limit int) ([]models.ApiPostDetailAndScore, error) {
	//获取排序后的最新 或最高分postid数组
	datas, dataswitchscore, err := daoredis.GetPostListCommunityKeyvalueByParam(communityid, typeId, cpage, limit)
	//如果无相关帖子可显示 则不继续下面的查询
	if len(datas) == 0 {
		return nil, err
	}
	//获取本次帖子相关的redis正数
	dataforvotesum, err := daoredis.GetPostListVoteNumByPostList(datas)
	//fmt.Println(dataforvotesum)
	//到mysql中获取帖子详情
	datasdetaillist, err := daomysql.GetPostListPostIdList(datas)
	//初始化结构体数组
	var returninfo = make([]models.ApiPostDetailAndScore, len(dataswitchscore))

	//排序输出
	for key, valueZ := range dataswitchscore {
		for _, value := range datasdetaillist {
			if valueZ.Member.(string) == strconv.FormatInt(value.Post_id, 10) {
				//获取当前帖子的社区信息
				datacommunity, err := daomysql.GetCommunityByCid(value.Community_id)
				if err != nil {
					return nil, err
				}
				//获取当前帖子的用户信息
				datauser, err := daomysql.GetUserInfoByUserId(value.Author_id)
				if err != nil {
					return nil, err
				}
				//为当前帖子的正数票数赋值
				var votenumtemp int64
				for votsumpostid, votsumvalue := range dataforvotesum {
					if votsumpostid == valueZ.Member.(string) {
						votenumtemp = votsumvalue
					}
				}

				returninfo[key] = models.ApiPostDetailAndScore{
					AuthorName:      datauser.Username,
					VoteNum:         votenumtemp,         //帖子投票为正数赞成票的计数
					Score:           int64(valueZ.Score), //时间为检索的时候为时间戳，分数检索的时候代表分数
					PostDetail:      value,
					CommunityDetail: *datacommunity,
				}
				break
			}
		}
	}

	return returninfo, err
}

//帖子详情
//使用了接口专用的模型ApiPostDetail
func GetPostDetailByPid(pid int64) (*models.ApiPostDetail, error) {
	//帖子详情数据使用ApiPostDetail结构体需要拼接
	datapost, err := daomysql.GetPostDetailByPid(pid)
	if err != nil {
		return nil, err
	}
	datacommunity, err := daomysql.GetCommunityByCid(datapost.Community_id)
	if err != nil {
		return nil, err
	}
	datauser, err := daomysql.GetUserInfoByUserId(datapost.Author_id)
	if err != nil {
		return nil, err
	}

	datas := &models.ApiPostDetail{
		AuthorName:      datauser.Username,
		PostDetail:      datapost,
		CommunityDetail: datacommunity,
	}
	return datas, err
}

//发布
func PostInfo(p *models.ParamUserPost, user_id int64) (err error) {
	//生成uid
	post_id, err := snowflake.GenId()

	postinfo := map[string]interface{}{
		"author_id":    user_id,
		"post_id":      post_id,
		"title":        p.Title,
		"content":      p.Content,
		"community_id": p.Community_id,
		"status":       1}
	err = daomysql.InsertPost(postinfo)
	if err == nil {
		//将发帖时间存入redis 作为是否可参与投票的判断时间依据
		daoredis.SavePostTimeAndInitScore(strconv.Itoa(int(post_id)))
		//将postid存入社区idkey set 作为后期社区内排序依据
		daoredis.SaveCommunityPostIdInSet(strconv.Itoa(int(p.Community_id)), strconv.Itoa(int(post_id)))
	}
	return err
}
