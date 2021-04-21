package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// StatusOK 運行成功
func StatusOK(c *gin.Context) (int, *Response) {
	return http.StatusOK, newResponse(c)
}

// StatusCreated 建立成功
func StatusCreated(c *gin.Context) (int, *Response) {
	return http.StatusCreated, newResponse(c)

}

//StatusAccepted 允許存取
func StatusAccepted(c *gin.Context) (int, *Response) {
	return http.StatusAccepted, newResponse(c)

}

//StatusNoContent 沒有資料
func StatusNoContent(c *gin.Context) (int, *Response) {
	return http.StatusNoContent, newResponse(c)
}

//SetResponse 設定 response 相關資料進入 gin 的 context
func SetResponse(c *gin.Context, message string, data interface{}, fails []string, pageInfo *PageInfo) {
	c.Set(MessageKey, message)
	c.Set(DataKey, data)
	if pageInfo != nil {
		c.Set(PageKey, pageInfo.Setup(c))
	}
	if fails != nil {
		c.Set(FailKey, fails)
	}
}
