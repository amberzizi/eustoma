package models

//统一使用的常量名称
const ContextUserIdKey = "user_id" //用户登录后上下文携带的user_id

//完整用户信息
type User struct {
	Id          int64 `json:"id,string"`
	User_id     int64 `json:"user_id,string"`
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
	Id       int64 `json:"id,string"`
	User_id  int64 `json:"user_id,string"`
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
	Id       int64 `json:"id,string"`
	User_id  int64 `json:"user_id,string"`
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
	User_id  int64 `json:"user_id,string"`
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
