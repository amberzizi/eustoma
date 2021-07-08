package daomysql

import "mygin/application/models"

//获取社区信息列表
func GetCommunitylist() ([]map[string]interface{}, error) {
	var connection = ReturnMsqlGoroseConnection()
	//var communitys models.Community
	db := connection.NewSession()
	info, err := db.Table("community").Get()
	return info, err
}

//根据社区id获取社区信息
func GetCommunityByCid(cid int64) (*models.Community, error) {
	var connection = ReturnMsqlGoroseConnection()
	var communitys models.Community
	db := connection.NewSession()
	err := db.Table(&communitys).Where("community_id", cid).Select()
	return &communitys, err
}
