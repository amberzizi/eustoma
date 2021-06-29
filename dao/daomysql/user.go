package daomysql


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
func InsertUser( user map[string]interface{}) (err error) {
	var connection = ReturnMsqlGoroseConnection()
	db := connection.NewSession()
	_ , err = db.Table("user").Data(user).Insert()
	return err
}
