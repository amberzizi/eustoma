package daomysql

import (
	"mygin/application/models"
)

//插重
func GetUserInfoByUsername(username string) (bool, error) {
	var connection = ReturnMsqlGoroseConnection()
	db := connection.NewSession()
	//var user models.User
	count, err := db.Table("user").Where("username", username).Count("*")
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//插入用户信息
func InsertUser(user map[string]interface{}) (err error) {
	var connection = ReturnMsqlGoroseConnection()
	db := connection.NewSession()
	_, err = db.Table("user").Data(user).Insert()
	return err
}

//获取用户信息 根据 user_id
func GetUserInfoByUserId(user_id int) (*models.Userinfopublic, error) {
	var connection = ReturnMsqlGoroseConnection()
	var userinfo models.Userinfopublic
	db := connection.NewSession()
	err := db.Table(&userinfo).Where("user_id", user_id).Select()
	return &userinfo, err
}

//获取用户鉴别登录信息 根据 username
func GetUserInfoByUsernameForLogin(username string) (*models.Userforlogin, error) {
	var connection = ReturnMsqlGoroseConnection()
	var userinfo models.Userforlogin
	db := connection.NewSession()
	err := db.Table(&userinfo).Where("username", username).Select()

	return &userinfo, err
}
