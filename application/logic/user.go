package logic

import (
	"errors"
	"mygin/application/models"
	"mygin/dao/daomysql"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//查重
	hads, err := daomysql.GetUserInfoByUsername(p.Username)
	if err != nil {
		return err
	}
	if hads {
		return errors.New("用户已存在")
	}

	println("不存在可生成")
	//生成uid
	//userID := snowflake.GenId()
	//盐
	//salt := randstring.RandSeq(10)
	//加盐后密码
	//finalpw := encryption.Md5(p.Password + salt)
	//密码加密

	return nil

}
