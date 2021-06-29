//用户 逻辑
package logic

import (
	"errors"
	"mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/tools/encryption"
	"mygin/tools/randstring"
	"mygin/tools/snowflake"
)

//用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	//查重
	hads, err := daomysql.GetUserInfoByUsername(p.Username)
	if err != nil {
		//如有执行报错 传递
		return err
	}

	if hads {
		//是否已有用户名相同用户  抛出错误
		return errors.New("用户已存在")
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
	userinfo, err := daomysql.GetUserInfoByUserId(p.User_id)
	return userinfo, err

}
