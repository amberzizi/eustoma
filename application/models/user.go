package models

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

func (u User) TableName() string {
	return "user"
}
