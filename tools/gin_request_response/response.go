package gin_request_response

import (
	"github.com/gin-gonic/gin"
	"mygin/settings"
	"net/http"
)

/**

 */

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//返回报错
//c.JSON(http.StatusOK, gin.H{
//	"code": settings.CodeInvalidParam,
//	"msg":  settings.CodeSetting[settings.CodeInvalidParam],
//	"data": [1]string{},
//})
//返回
func Response(c *gin.Context, code int, data interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  settings.CodeSetting[code],
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func Response401(c *gin.Context, code int, data interface{}) {
	rd := &ResponseData{
		Code: code,
		Msg:  settings.CodeSetting[code],
		Data: data,
	}
	c.JSON(http.StatusUnauthorized, rd)
}
