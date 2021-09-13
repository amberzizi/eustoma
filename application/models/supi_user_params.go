package models

//SUPI 用户 提交参数

//小程序用户注册
type ParamUserMiniappSignUp struct {
	Openid         string `json:"openid" binding:"required"`
	Unionid        string `json:"unionid"`
	Belong_miniapp string `json:"belong_miniapp" binding:"required"`
	Share_uid      string `json:"share_uid"`
}
