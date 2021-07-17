package daomysql

import (
	"mygin/application/models"
)

func GetPostListPostIdList(datas []string) ([]models.PostDetail, error) {
	db := gdb.NewSession()
	var posts []models.PostDetail
	err := db.Table(&posts).WhereIn("post_id", datas).Limit(len(datas)).Select()
	return posts, err
}
