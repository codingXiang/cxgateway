package e

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
func StatusInternalServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, SERVER_ERROR, StatusText(SERVER_ERROR))
}

//StatusNotImplemented 501
func StatusNotImplemented(message string) *APIException {
	return newAPIException(http.StatusNotImplemented, SERVER_ERROR, message)
}

//StatusBadGateway 502
func StatusBadGateway(message string) *APIException {
	return newAPIException(http.StatusBadGateway, SERVER_ERROR, message)
}

//StatusServiceUnavailable 503
func StatusServiceUnavailable(message string) *APIException {
	return newAPIException(http.StatusServiceUnavailable, SERVER_ERROR, message)
}

//StatusGatewayTimeout 504
func StatusGatewayTimeout(message string) *APIException {
	return newAPIException(http.StatusGatewayTimeout, SERVER_ERROR, message)
}

/*
 400 系列
*/
//StatusBadRequest 400
func StatusBadRequest(message string) *APIException {
	return newAPIException(http.StatusBadRequest, PARAMETER_ERROR, message)
}

//StatusUnauthorized 401
func StatusUnauthorized(message string) *APIException {
	return newAPIException(http.StatusUnauthorized, AUTH_ERROR, message)
}

//StatusForbidden 403
func StatusForbidden(message string) *APIException {
	return newAPIException(http.StatusForbidden, AUTH_ERROR, message)
}

//NotFoundError 404 錯誤
func StatusNotFound(message string) *APIException {
	return newAPIException(http.StatusNotFound, NOT_FOUND, message)
}

//StatusConflict 409
func StatusConflict(message string) *APIException {
	return newAPIException(http.StatusConflict, DUPLICATE_ERROR, message)
}

//UnknownError 未知錯誤
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}


//HandleNotFound 處理404頁面
func HandleNotFound(c *gin.Context) {
	handleErr := StatusNotFound("this api is not found")
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}

//Wrapper 在 register routing 時加入錯誤 handler
func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == gin.DebugMode {
					apiException = UnknownError(e.Error())
				} else {
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = StatusInternalServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}
