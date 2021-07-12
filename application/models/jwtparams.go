package models

import (
	"github.com/dgrijalva/jwt-go"
	"mygin/settings"
	"time"
)

//定义有效时间为2小时 accesstoken有效时间
const TokenExpireDuration = time.Second * 60 * 10 //accesstoken 十分钟过期
//const TokenExpireDuration = time.Second * 60 //accesstoken 1分钟过期  测试
const RefreshTokenExpireDuration = time.Hour * 24 * 7 //刷新token过期时间较长
//const RefreshTokenExpireDuration = time.Second * 600 //刷新token过期时间较长 测试

//jwt 携带的数据结构体
//jwt包自带jwt.standandclaims 只包含了官方字段  需额外字段需自定义结构体
type MyJwtInfo struct {
	Username string `json:"username"`
	User_id  int64  `json:"user_id,string"`
	jwt.StandardClaims
}

//获取jwt签名加密key
func GetJwtTokenKey() interface{} {
	return []byte(settings.SettingGlb.App.Jwtkey)
}
