package models

import (
	"github.com/dgrijalva/jwt-go"
	"mygin/settings"
	"time"
)

//定义有效时间为2小时
const TokenExpireDuration = time.Hour * 2

var JwtTokenKey *[]byte

//jwt 携带的数据结构体
//jwt包自带jwt.standandclaims 只包含了官方字段  需额外字段需自定义结构体
type MyJwtInfo struct {
	Username string `json:"username"`
	User_id  int64  `json:"user_id"`
	jwt.StandardClaims
}

func GetJwtTokenKey() interface{} {
	return []byte(settings.SettingGlb.App.Jwtkey)
}
