package util

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestHandlerInterface interface {
	//BindBody : 綁定 body
	BindBody(c *gin.Context, body interface{}) error
	//ValidValidation : 驗證表單資訊
	ValidValidation(c *gin.Context, v *validation.Validation) error
}

type RequestHandler struct {
	context *gin.Context
}

func NewRequestHandler() RequestHandlerInterface {
	return &RequestHandler{}
}

//BindBody : 綁定 body
func (r *RequestHandler) BindBody(c *gin.Context, body interface{}) error {
	var err = c.Bind(&body)
	if err != nil {
		return response.StatusBadRequest(c)
	}
	return nil
}

//ValidValidation : 驗證表單資訊
func (r *RequestHandler) ValidValidation(c *gin.Context, v *validation.Validation) error {
	if v.HasErrors() {
		var fails []string = make([]string, 0)
		for _, err := range v.Errors {
			fails = append(fails, err.Error())
		}
		return response.StatusBadRequest(c)
	}
	return nil
}

type RequesterInterface interface {
	GET(uri string) (*response.Response, error)
	POST(uri string, param interface{}) (*response.Response, error)
	PUT(uri string, param interface{}) (*response.Response, error)
	PATCH(uri string, param interface{}) (*response.Response, error)
	DELETE(uri string, param interface{}) (*response.Response, error)
	GetWithHeader(uri string, header map[string]string) (*response.Response, error)
	PostWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error)
	PutWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error)
	PatchWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error)
	DeleteWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error)
}

type Requester struct {
	client *http.Client
}

func NewRequester(client *http.Client) RequesterInterface {
	if client == nil {
		client = &http.Client{}
	}
	return &Requester{
		client: client,
	}
}

func (r *Requester) ReadJSONResponse(in *http.Response) (*response.Response, error) {
	defer in.Body.Close()
	var resp = new(response.Response)
	err := json.NewDecoder(in.Body).Decode(resp)
	return resp, err
}

func (r *Requester) GET(uri string) (*response.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
func (r *Requester) POST(uri string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}

func (r *Requester) PUT(uri string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}

func (r *Requester) PATCH(uri string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}

func (r *Requester) DELETE(uri string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)

	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodDelete, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}

func (r *Requester) GetWithHeader(uri string, header map[string]string) (*response.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
func (r *Requester) PostWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)
	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
func (r *Requester) DeleteWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)
	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodDelete, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
func (r *Requester) PutWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)
	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
func (r *Requester) PatchWithHeader(uri string, header map[string]string, param interface{}) (*response.Response, error) {
	// 轉換參數
	jsonByte, _ := json.Marshal(param)
	// 將 json data 放在 body 進行 request
	req, err := http.NewRequest(http.MethodPatch, uri, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	return r.ReadJSONResponse(resp)
}
