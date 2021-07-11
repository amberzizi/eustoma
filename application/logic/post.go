package logic

import (
	"mygin/application/models"
	"mygin/dao/daomysql"
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
