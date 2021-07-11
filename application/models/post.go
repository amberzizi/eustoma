package models

//分类模型
//完整的分类信息
type Post struct {
	//Id             int64
	Post_id   int64
	Title     string
	Author_id int64
	//Introduction   string
	//Create_time    string
	//Update_time    string
}

type PostDetail struct {
	Id           int64
	Post_id      int64
	Title        string
	Content      string
	Author_id    int64
	Community_id int64
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
