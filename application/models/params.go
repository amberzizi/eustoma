package models

//用户注册传入参数模型
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required,email"`
	Age        int    `json:"age" binding:"gte=1,lte=130"`
}
