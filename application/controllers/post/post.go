package post

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygin/application/logic"
	"mygin/application/models"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"strconv"
)

//帖子列表
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

//帖子详情
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

//发布帖子
func PostInfoHandle(c *gin.Context) {
	//校验参数
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamUserPost)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("PostInfoHandle with invalid param", zap.Error(err))
		//返回
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//获取发布者user_id
	user_id, err := gin_request_response.GetCurrentUser(c)
	if err != nil {
		gin_request_response.Response(c, settings.ErrorUserNotLogin, nil)
	}

	err = logic.PostInfo(p, user_id)
	if err != nil {
		gin_request_response.Response(c, settings.CodePostError, nil)
	}
	gin_request_response.Response(c, settings.CodeSuccess, nil)
}

//帖子投票
func PostVoteHandle(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamVotePost)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("PostInfoHandle with invalid param", zap.Error(err))
		//返回
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//获取发布者user_id
	user_id, err := gin_request_response.GetCurrentUser(c)
	if err != nil {
		gin_request_response.Response(c, settings.ErrorUserNotLogin, nil)
	}

	result, err := logic.PostVote(p, user_id)
	if errors.Is(err, models.ErrorVoteOutOfTime) {
		gin_request_response.Response(c, settings.ErrorVoteOutOfTime, nil)
		return
	}
	if err != nil || !result {
		gin_request_response.Response(c, settings.CodeVoteError, nil)
	}
	gin_request_response.Response(c, settings.CodeSuccess, nil)
}
