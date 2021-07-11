package logic

import (
	"mygin/application/models"
	"mygin/dao/daomysql"
)

//社区列表
func GetCommunityList() ([]map[string]interface{}, error) {
	datas, err := daomysql.GetCommunitylist()
	return datas, err
}

//社区信息
func GetCommunityInfoById(id int64) (*models.CommunityDetail, error) {
	datas, err := daomysql.GetCommunityByCid(id)
	return datas, err
}
