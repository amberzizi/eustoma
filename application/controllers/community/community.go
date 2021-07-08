package community

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygin/application/logic"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"strconv"
)

//获取社区分类列表
func CommunityHandle(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("CommunityHandle faild", zap.Error(err))
		gin_request_response.Response(c, settings.CodeServerBusy, nil)
		return
	}
	gin_request_response.Response(c, settings.CodeSuccess, data)
}

//根据id获取社区分类信息
func GetCommunityInfoById(c *gin.Context) {
	communityId := c.Param("communityId")
	cid, err := strconv.ParseInt(communityId, 10, 64)
	data, err := logic.GetCommunityInfoById(cid)
	if err != nil {
		zap.L().Error("GetCommunityInfoById faild", zap.Error(err))
		gin_request_response.Response(c, settings.CodeServerBusy, nil)
		return
	}
	gin_request_response.Response(c, settings.CodeSuccess, data)
}
