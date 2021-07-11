package models

//分类模型
//完整的分类信息
type Community struct {
	//Id             int64
	Community_id   int64
	Community_name string
	//Introduction   string
	//Create_time    string
	//Update_time    string
}

type CommunityDetail struct {
	Id             int64
	Community_id   int64
	Community_name string
	Introduction   string
	Create_time    string
	Update_time    string
}

func (u Community) TableName() string {
	return "community"
}

func (u CommunityDetail) TableName() string {
	return "community"
}
