package daoredis

//查询
func GetAccesstokenByUserID(user_id string) (string, error) {
	rdb := ReturnRedisDb()
	n, err := rdb.Get(user_id).Result()
	return n, err
}

//设置
func SetAccesstokenByUserID(user_id string, atoken string) error {
	rdb := ReturnRedisDb()
	err := rdb.Set(user_id, atoken, 0).Err()
	return err
}
