package e

import "net/http"

// 運行成功
func StatusOK(message string, data interface{}) (int, *Response) {
	return http.StatusOK, newResponse(SUCCESS, message, data)
}

// 建立成功
func StatusCreated(message string, data interface{}) (int, *Response) {
	return http.StatusCreated, newResponse(CREATED, message, data)
}

func StatusAccepted(message string, data interface{}) (int, *Response) {
	return http.StatusAccepted, newResponse(ACCEPT, message, data)
}

//StatusNoContent 沒有資料
func StatusNoContent(message string) (int, *Response) {
	return http.StatusNoContent, newResponse(ACCEPT, message, nil)
}
