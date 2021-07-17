package daomysql

import "mygin/application/models"

func GetPostListByCid(cid int64, page int, limit int) ([]models.PostDetail, error) {
	var connection = ReturnMsqlGoroseConnection()
	var posts []models.PostDetail
	db := connection.NewSession()
	currentoffset := 0
	if page > 1 {
		currentoffset = page*limit - 1
	}
	err := db.Table(&posts).Where("community_id", cid).Where("status", 1).OrderBy("id desc").Limit(limit).Offset(currentoffset).Select()
	return posts, err
}

func GetPostDetailByPid(pid int64) (*models.PostDetail, error) {
	var connection = ReturnMsqlGoroseConnection()
	var postinfo models.PostDetail
	db := connection.NewSession()
	err := db.Table(&postinfo).Where("post_id", pid).Select()
	return &postinfo, err
}

func InsertPost(posinfo map[string]interface{}) (err error) {
	var connection = ReturnMsqlGoroseConnection()
	db := connection.NewSession()
	_, err = db.Table("post").Data(posinfo).Insert()
	return err
}

//获取
