package gin_request_response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mygin/application/models"
)

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (user_id int64, err error) {
	uid, exsit := c.Get(models.ContextUserIdKey)
	if !exsit {
		err = ErrorUserNotLogin
		return
	}
	user_id, ok := uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return user_id, err
}
