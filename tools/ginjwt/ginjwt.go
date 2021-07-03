package ginjwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mygin/application/models"
	"time"
)

//无效token 返回error
var (
	ErrorInvalidToken = errors.New("Invalid token")
)

//生成token
func GenJwtToken(username string, user_id int64) (string, error) {
	//创建playload
	c := models.MyJwtInfo{
		username,
		user_id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(models.TokenExpireDuration).Unix(),
			Issuer:    "eustoma",
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secretkey签名并获得完整的编码后的字符串token
	return token.SignedString(models.GetJwtTokenKey())
}

//解析token
func ParseJwtToken(tokenString string) (*models.MyJwtInfo, error) {
	var mc = new(models.MyJwtInfo)
	token, err := jwt.ParseWithClaims(
		tokenString,
		mc,
		func(token *jwt.Token) (i interface{}, err error) {
			return models.GetJwtTokenKey(), nil
		})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, ErrorInvalidToken
}
