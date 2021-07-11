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
func GetPostDetailByPid(pid int64) (*models.PostDetail, error) {
	datas, err := daomysql.GetPostDetailByPid(pid)
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
