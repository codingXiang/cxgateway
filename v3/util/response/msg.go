package response

import "github.com/gin-gonic/gin"

type PageInfo struct {
	Count       int `json:"count"`       //全部資料數量
	Limit       int `json:"limit"`       //限制搜尋筆數
	TotalPage   int `json:"totalPage"`   //總共頁數
	CurrentPage int `json:"currentPage"` //目前所在頁數
}

type Response struct {
	Code     Code        `json:"code"`               //系統定義代碼
	Message  string      `json:"message"`            //訊息
	Data     interface{} `json:"data"`               //回傳資料
	PageInfo *PageInfo   `json:"pageInfo,omitempty"` //分頁資訊
	Fails    []string    `json:"fails,omitempty"`    //錯誤訊息
}

// code 專用
const (
	CodeStatus = "codeStatus"
	ModuleName = "moduleName"
	MethodName = "methodName"
)

// message
const (
	MessageKey = "responseMessage"
)

//pageInfo
const (
	PageKey = "pageInfo"
)

//data
const (
	DataKey = "responseData"
)

//fail
const (
	FailKey = "responseFails"
)

func newResponse(c *gin.Context) *Response {
	var (
		pageInfo *PageInfo   = nil
		data     interface{} = nil
	)

	if page, exist := c.Get(PageKey); exist {
		pageInfo = page.(*PageInfo)
	}

	if in, exist := c.Get(DataKey); exist {
		data = in
	}

	return &Response{
		Code:     NewCode(c),
		Message:  c.GetString(MessageKey),
		Data:     data,
		PageInfo: pageInfo,
		Fails:    c.GetStringSlice(FailKey),
	}
}

func (r *Response) Error() string {
	return r.Message
}

//type JsonResponse struct {
//	StatusCode int
//	*Response
//}
//
//func NewJsonResponse(statusCode int, response *Response) (int, *JsonResponse) {
//	return &JsonResponse{
//		StatusCode: statusCode,
//		Response:   response,
//	}
//}
//
//func (response *JsonResponse) Error() string {
//	return response.Message
//}

//const (
//	SERVER_ERROR    = 10000 //系統錯誤
//	NOT_FOUND       = 10001 //找不到頁面
//	UNKNOWN_ERROR   = 10002 //未知的錯誤
//	PARAMETER_ERROR = 10003 //參數錯誤
//	AUTH_ERROR      = 10004 //驗證錯誤
//	NO_CONTENT      = 10005 //沒有內容
//	DUPLICATE_ERROR = 10006 //重複資料
//	SUCCESS         = 20000 //運行成功
//	CREATED         = 20001 //建立成功
//	ACCEPT          = 20002 //允許操作
//)
//
//var statusText = map[int]string{
//	SERVER_ERROR:    "伺服器發生錯誤，請通知相關人員",
//	NOT_FOUND:       "找不到此頁面",
//	UNKNOWN_ERROR:   "未知的錯誤",
//	PARAMETER_ERROR: "參數輸入錯誤",
//	AUTH_ERROR:      "token驗證錯誤",
//	NO_CONTENT:      "找不到內容",
//	SUCCESS:         "成功",
//	CREATED:         "建立完成",
//}
//
//func StatusText(code int) string {
//	return statusText[code]
//}
