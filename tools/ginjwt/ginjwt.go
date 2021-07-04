package ginjwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"mygin/application/models"
	"time"
)

//无效token 返回error
var (
	ErrorInvalidToken                = errors.New("Invalid token")
	ErrorAccessTokenExpiredOutOfTime = errors.New("accesstoken过期")
)

//生成accesstoken
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

//生成refreshtoken
func GenJwtRefreshToken() (string, error) {
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(models.RefreshTokenExpireDuration).Unix(),
		Issuer:    "eustoma",
	})
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
		//捕获到期错误
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, ErrorAccessTokenExpiredOutOfTime
		}
		return nil, err
	}
	if token.Valid {
		return mc, err
	}
	return nil, ErrorInvalidToken
}

//使用freshtoken 更新 accesstoken
func RefreshTokenForNewAccessToken(aToken, rToken string) (string, error) {
	//refresh token 无效直接返回
	if _, err := jwt.Parse(rToken,
		func(token *jwt.Token) (i interface{}, err error) {
			return models.GetJwtTokenKey(), nil
		}); err != nil {
		return "", err
	}
	//从旧的accesstoken中解析出claims数据
	var claims models.MyJwtInfo
	_, err := jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (i interface{}, err error) {
		return models.GetJwtTokenKey(), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		//当accesstoken 是过期错误 并且refreshtoken没过期时就创建并返回一个新的accesstoken
		if v.Errors == jwt.ValidationErrorExpired {
			return GenJwtToken(claims.Username, claims.User_id)
		}
	}

	//如果解析未出问题 原始accesstoken也未过期 返回原始accesstoken
	return aToken, err
}
