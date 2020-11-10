package e

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	Request   string `json:"request"`
}

const (
	SERVER_ERROR    = 10000 //系統錯誤
	NOT_FOUND       = 10001 //找不到頁面
	UNKNOWN_ERROR   = 10002 //未知的錯誤
	PARAMETER_ERROR = 10003 //參數錯誤
	AUTH_ERROR      = 10004 //驗證錯誤
	NO_CONTENT      = 10005 //沒有內容
	DUPLICATE_ERROR = 10006 //重複資料
	SUCCESS         = 20000 //運行成功
	CREATED         = 20001 //建立成功
	ACCEPT          = 20002 //允許操作
)

var statusText = map[int]string{
	SERVER_ERROR:    "伺服器發生錯誤，請通知相關人員",
	NOT_FOUND:       "找不到此頁面",
	UNKNOWN_ERROR:   "未知的錯誤",
	PARAMETER_ERROR: "參數輸入錯誤",
	AUTH_ERROR:      "token驗證錯誤",
	NO_CONTENT:      "找不到內容",
	SUCCESS:         "成功",
	CREATED:         "建立完成",
}

func StatusText(code int) string {
	return statusText[code]
}

func (e *APIException) Error() string {
	return e.Message
}

func newAPIException(code int, errorCode int, message string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Message:   message,
	}
}

func newResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
