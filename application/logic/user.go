//用户 逻辑
package logic

import (
	"database/sql"
	"mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/dao/daoredis"
	"mygin/tools/encryption"
	"mygin/tools/ginjwt"
	"mygin/tools/randstring"
	"mygin/tools/snowflake"
	"strconv"
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
		return models.ErrorUserExist
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
func LoginCheckPassword(p *models.ParamLoginIn) (bool, error, *models.Userinfopublic) {
	userinfo, err := daomysql.GetUserInfoByUsernameForLogin(p.Username)
	if err == sql.ErrNoRows {
		return false, models.ErrorUserNotExist, nil
	}
	if err != nil {
		return false, err, nil
	}
	aftermd5pw, err := encryption.Md5(p.Password + userinfo.Salt)
	if err != nil {
		return false, err, nil
	}
	if aftermd5pw == userinfo.Password {
		userinfopublic, err := daomysql.GetUserInfoByUserId(userinfo.User_id)
		return true, err, userinfopublic
	}
	return false, models.ErrorUserPassowrdWrong, nil
}

//生成jwttoken accesstoken
func GenUserJwtToken(p *models.ParamLoginIn) (string, error) {
	userinfo, err := daomysql.GetUserInfoByUsernameForJWT(p.Username)
	if err == sql.ErrNoRows {
		return "", models.ErrorUserNotExist
	}
	if err != nil {
		return "", err
	}
	token, err := ginjwt.GenJwtToken(p.Username, userinfo.User_id)
	if err != nil {
		return "", err
	}
	//成功登录 ，保存对应关系到redis 用以单点登录检查
	SetAccesstokenByUserID_ForRelation(userinfo.User_id, token)

	return token, err
}

//生成jwttoken refreshtoken
func GenUserJwtRefreshToken() (string, error) {
	token, err := ginjwt.GenJwtRefreshToken()
	if err != nil {
		return "", err
	}
	return token, err
}

//刷新accesstoken
func GetUserNewAccesstoken(p *models.ParamRefreshAccessToken) (string, error) {
	token, user_id, err := ginjwt.RefreshTokenForNewAccessToken(p.Accesstoken, p.Refreshtoken)
	if err != nil {
		return "", err
	}
	//成功刷新 刷新accesstoken后 重新为userid绑定
	SetAccesstokenByUserID_ForRelation(user_id, token)
	return token, err
}

//从数据库内取出userid对应的accecctoken redis
func GetOnlineAccesstokenByUserID_ForRelation(user_id int64) (string, error) {
	re, err := daoredis.GetAccesstokenByUserID("single_login_AT_" + strconv.FormatInt(user_id, 10))
	return re, err
}

//设置对应accesstoken  redis
func SetAccesstokenByUserID_ForRelation(user_id int64, atoken string) error {
	err := daoredis.SetAccesstokenByUserID("single_login_AT_"+strconv.FormatInt(user_id, 10), atoken)
	return err
}

////解析jwttoken
//func ParseUserJwtToken(tokenString string) (*models.Userinfopublic, error) {
//	jwtinfo, err := ginjwt.ParseJwtToken(tokenString)
//	if err != nil {
//		return nil, err
//	}
//
//	userinfo, err := daomysql.GetUserInfoByUserId(jwtinfo.User_id)
//	if err != nil {
//		return nil, err
//	}
//
//	return userinfo, err
//}
