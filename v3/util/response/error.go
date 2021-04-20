package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerFunc 錯誤 handler
type HandlerFunc func(c *gin.Context) error

/*
	500 系列
*/
//ServerError 500 伺服器錯誤
func StatusInternalServerError(c *gin.Context) (int, *Response) {
	return http.StatusInternalServerError, newResponse(c)
}

//StatusNotImplemented 501
func StatusNotImplemented(c *gin.Context) (int, *Response) {
	return http.StatusNotImplemented, newResponse(c)
}

//StatusBadGateway 502
func StatusBadGateway(c *gin.Context) (int, *Response) {
	return http.StatusBadGateway, newResponse(c)
}

//StatusServiceUnavailable 503
func StatusServiceUnavailable(c *gin.Context) (int, *Response) {
	return http.StatusServiceUnavailable, newResponse(c)
}

//StatusGatewayTimeout 504
func StatusGatewayTimeout(c *gin.Context) (int, *Response) {
	return http.StatusGatewayTimeout, newResponse(c)
}

/*
 400 系列
*/
//StatusBadRequest 400
func StatusBadRequest(c *gin.Context) (int, *Response) {
	return http.StatusBadRequest, newResponse(c)
}

//StatusUnauthorized 401
func StatusUnauthorized(c *gin.Context) (int, *Response) {
	return http.StatusUnauthorized, newResponse(c)
}

//StatusForbidden 403
func StatusForbidden(c *gin.Context) (int, *Response) {
	return http.StatusForbidden, newResponse(c)
}

//NotFoundError 404 錯誤
func StatusNotFound(c *gin.Context) (int, *Response) {
	return http.StatusNotFound, newResponse(c)
}

//StatusConflict 409
func StatusConflict(c *gin.Context) (int, *Response) {
	return http.StatusConflict, newResponse(c)
}

//UnknownError 未知錯誤，回傳 500
func UnknownError(c *gin.Context) (int, *Response) {
	return http.StatusInternalServerError, newResponse(c)
}

//Wrapper 在 register routing 時加入錯誤 handler
func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *Response
			if h, ok := err.(*Response); ok {
				apiException = h
			} else if _, ok := err.(error); ok {
				_, apiException = UnknownError(c)
			} else {
				c.JSON(StatusInternalServerError(c))
				return
			}
			c.Set(MethodName, c.Request.Method)
			c.Set(DataKey, apiException)
			//apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(StatusInternalServerError(c))
			return
		}
	}
}
