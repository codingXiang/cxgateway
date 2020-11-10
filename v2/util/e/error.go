package e

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerFunc 錯誤 handler
type HandlerFunc func(c *gin.Context) error

//ServerError 500 伺服器錯誤
func ServerError() *APIException {
	var err = SERVER_ERROR
	return newAPIException(http.StatusInternalServerError, err, StatusText(err))
}

//NotFoundError 404 錯誤
func NotFoundError(message string) *APIException {
	var err = NOT_FOUND
	return newAPIException(http.StatusNotFound, err, message)
}

//UnknownError 未知錯誤
func UnknownError(message string) *APIException {
	var err = UNKNOWN_ERROR
	return newAPIException(http.StatusUnprocessableEntity, err, message)
}

//ParameterError 參數錯誤
func ParameterError(message string) *APIException {
	var err = PARAMETER_ERROR
	return newAPIException(http.StatusBadRequest, err, message)
}

//DuplicateError 重複資料
func DuplicateError(message string) *APIException {
	var err = DUPLICATE_ERROR
	return newAPIException(http.StatusConflict, err, message)
}

//NoContentError 沒有資料
func NoContentError(message string) *APIException {
	var err = NO_CONTENT
	return newAPIException(http.StatusNoContent, err, message)
}

//AuthError token 驗證錯誤
func AuthError(message string) *APIException {
	var err = AUTH_ERROR
	return newAPIException(http.StatusUnauthorized, err, message)
}

//SuccessError 未知錯誤
func SuccessError(message string) *APIException {
	var err = SUCCESS
	return newAPIException(http.StatusOK, err, message)
}

//HandleNotFound 處理404頁面
func HandleNotFound(c *gin.Context) {
	handleErr := NotFoundError("this api is not found")
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
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			apiException.Request = c.Request.Method + " " + c.Request.URL.String()
			c.JSON(apiException.Code, apiException)
			return
		}
	}
}
