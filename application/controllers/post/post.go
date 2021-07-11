package post

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygin/application/logic"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"strconv"
)

func PostListHandle(c *gin.Context) {
	communityId := c.Param("communityId")
	page := c.Param("page")
	limit := c.Param("limit")
	cid, errc := strconv.ParseInt(communityId, 10, 64)
	cpage, errp := strconv.Atoi(page)
	climit, errl := strconv.Atoi(limit)
	if errc != nil || errp != nil || errl != nil {
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}
	if climit > 99 {
		//每页面显示数量过多
		gin_request_response.Response(c, settings.CodeOutOfRange, nil)
		return
	}
	data, err := logic.GetPostListByCid(cid, cpage, climit)
	if err != nil {
		zap.L().Error("PostListHandle faild", zap.Error(err))
		gin_request_response.Response(c, settings.CodeServerBusy, nil)
		return
	}
	gin_request_response.Response(c, settings.CodeSuccess, data)
}

func PostListDetailHandle(c *gin.Context) {
	postId := c.Param("postId")
	pid, errc := strconv.ParseInt(postId, 10, 64)
	if errc != nil {
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}
	data, err := logic.GetPostDetailByPid(pid)
	if err != nil {
		zap.L().Error("PostListHandle faild", zap.Error(err))
		gin_request_response.Response(c, settings.CodeServerBusy, nil)
		return
	}
	gin_request_response.Response(c, settings.CodeSuccess, data)
}
