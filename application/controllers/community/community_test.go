package community

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommunityHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/community"
	r.POST(url, CommunityHandle)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	//判断响应内容是否包含指定字符串
	assert.Contains(t, w.Body.String(), "登录")

}
