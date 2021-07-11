package logic

import (
	"mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/tools/snowflake"
)

//社区列表
func GetPostListByCid(cid int64, page int, limit int) ([]map[string]interface{}, error) {
	datas, err := daomysql.GetPostListByCid(cid, page, limit)
	return datas, err
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
	return err
}
