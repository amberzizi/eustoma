package daoredis

//查询 用于单点登录 user_id =>token
func GetAccesstokenByUserID(user_id string) (string, error) {
	rdb := ReturnRedisDb()
	n, err := rdb.Get(user_id).Result()
	return n, err
}

//设置 用于单点登录
func SetAccesstokenByUserID(user_id string, atoken string) error {
	rdb := ReturnRedisDb()
	err := rdb.Set(user_id, atoken, 0).Err()
	return err
}
