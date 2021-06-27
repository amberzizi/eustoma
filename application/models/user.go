package myginuser

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) TableName() string {
	return "user"
}
