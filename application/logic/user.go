//用户 逻辑
package logic

import (
	"database/sql"
	"errors"
	"mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/tools/encryption"
	"mygin/tools/ginjwt"
	"mygin/tools/randstring"
	"mygin/tools/snowflake"
)

//定义自定义error
var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorUserNotExist      = errors.New("用户名不存在")
	ErrorUserPassowrdWrong = errors.New("用户密码错误")
)

//用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	//查重
	hads, err := daomysql.CheckUserInfoByUsername(p.Username)
	if err != nil {
		//如有执行报错 传递
		return err
	}

	if hads {
		//是否已有用户名相同用户  抛出错误
		return ErrorUserExist
	}

	//生成uid
	userID, err := snowflake.GenId()
	//盐
	salt := randstring.RandSeq(10)
	//加盐后密码
	finalpw, err := encryption.Md5(p.Password + salt)

	//return errors.New("用户已存在")
	userinfo := map[string]interface{}{
		"user_id":  userID,
		"username": p.Username,
		"password": finalpw,
		"salt":     salt,
		"email":    p.Email}
	err = daomysql.InsertUser(userinfo)

	return err
}

//获取用户信息 by userid
func GetUserInfoByUserId(p *models.ParamGetuserinfoByUID) (*models.Userinfopublic, error) {
	userinfo, err := daomysql.GetUserInfoByUserId(int64(p.User_id))
	return userinfo, err

}

//用户登录 校验密码
func LoginCheckPassword(p *models.ParamLoginIn) (bool, error) {
	userinfo, err := daomysql.GetUserInfoByUsernameForLogin(p.Username)
	if err == sql.ErrNoRows {
		return false, ErrorUserNotExist
	}
	if err != nil {
		return false, err
	}
	aftermd5pw, err := encryption.Md5(p.Password + userinfo.Salt)
	if err != nil {
		return false, err
	}
	if aftermd5pw == userinfo.Password {
		return true, nil
	}
	return false, ErrorUserPassowrdWrong
}

//生成jwttoken
func GenUserJwtToken(p *models.ParamLoginIn) (string, error) {
	userinfo, err := daomysql.GetUserInfoByUsernameForJWT(p.Username)
	if err == sql.ErrNoRows {
		return "", ErrorUserNotExist
	}
	if err != nil {
		return "", err
	}
	token, err := ginjwt.GenJwtToken(p.Username, userinfo.User_id)
	if err != nil {
		return "", err
	}
	return token, err
}

//解析jwttoken
func ParseUserJwtToken(tokenString string) (*models.Userinfopublic, error) {
	jwtinfo, err := ginjwt.ParseJwtToken(tokenString)
	if err != nil {
		return nil, err
	}

	userinfo, err := daomysql.GetUserInfoByUserId(jwtinfo.User_id)
	if err != nil {
		return nil, err
	}

	return userinfo, err
}
