package models

//完整用户信息
type User struct {
	Id          int64
	User_id     int64
	Username    string
	Password    string
	Salt        string
	Email       string
	Gender      int
	Create_time string
	Update_time string
}

//外部可用的用户信息
type Userinfopublic struct {
	Id       int64
	User_id  int64
	Username string
	//Password    string
	//Salt        string
	Email       string
	Gender      int
	Create_time string
	Update_time string
}

//用来登录鉴别的user信息
type Userforlogin struct {
	Id       int64
	User_id  int64
	Username string
	Password string
	Salt     string
	//Email       string
	//Gender      int
	//Create_time string
	//Update_time string
}

//用来生成jwt的信息
type Userforjwt struct {
	User_id  int64
	Username string
}

func (u User) TableName() string {
	return "user"
}

func (u Userinfopublic) TableName() string {
	return "user"
}

func (u Userforlogin) TableName() string {
	return "user"
}
func (u Userforjwt) TableName() string {
	return "user"
}
