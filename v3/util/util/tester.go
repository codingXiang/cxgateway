package util

import (
	"bytes"
	"encoding/json"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type (
	HttpTesterInterface interface {
		GET(uri string) (int, *response.Response)
		POST(uri string, param interface{}) (int, *response.Response)
		POST_FORM(uri string, param interface{}) (int, *response.Response)
		PUT(uri string, param interface{}) (int, *response.Response)
		PATCH(uri string, param interface{}) (int, *response.Response)
		DELETE(uri string, param interface{}) (int, *response.Response)
	}
	HttpTester struct {
		router *gin.Engine
	}
)

func NewHttpTester(router *gin.Engine) HttpTesterInterface {
	return &HttpTester{
		router: router,
	}
}

//Get 模擬Get Request的動作
func (h *HttpTester) GET(uri string) (int, *response.Response) {

	var (
		statusCode int
		response   = new(response.Response)
	)

	// 构造get请求
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)

	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}

//POST 模擬Post Request的動作
func (h *HttpTester) POST(uri string, param interface{}) (int, *response.Response) {
	var (
		statusCode int
		response   = new(response.Response)
	)
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req := httptest.NewRequest(http.MethodPost, uri, bytes.NewReader(jsonByte))

	// 初始化 rquest
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)
	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}

//POST_FORM 模擬Post Request的動作
func (h *HttpTester) POST_FORM(uri string, param interface{}) (int, *response.Response) {
	var (
		statusCode int
		response   = new(response.Response)
	)
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req := httptest.NewRequest(http.MethodPost, uri, bytes.NewReader(jsonByte))

	// 初始化 rquest
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)
	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}
//PUT 模擬Post Request的動作
func (h *HttpTester) PUT(uri string, param interface{}) (int, *response.Response) {
	var (
		statusCode int
		response   = new(response.Response)
	)
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req := httptest.NewRequest(http.MethodPut, uri, bytes.NewReader(jsonByte))

	// 初始化 rquest
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)
	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}

//PATCH 模擬 Patch Request的動作
func (h *HttpTester) PATCH(uri string, param interface{}) (int, *response.Response) {
	var (
		statusCode int
		response   = new(response.Response)
	)
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewReader(jsonByte))

	// 初始化 rquest
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)
	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}

//DELETE 模擬 Delete Request的動作
func (h *HttpTester) DELETE(uri string, param interface{}) (int, *response.Response) {
	var (
		statusCode int
		response   = new(response.Response)
	)
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req := httptest.NewRequest(http.MethodDelete, uri, bytes.NewReader(jsonByte))

	// 初始化 rquest
	w := httptest.NewRecorder()

	// 模擬調用
	h.router.ServeHTTP(w, req)

	// 取得 Response
	result := w.Result()
	defer result.Body.Close()

	// 讀取 Response body
	body, _ := ioutil.ReadAll(result.Body)
	statusCode = result.StatusCode

	json.Unmarshal(body, response)

	return statusCode, response
}
