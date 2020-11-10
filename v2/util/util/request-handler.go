package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/cxgateway/v2/util/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestHandlerInterface interface {
	//BindBody : 綁定 body
	BindBody(c *gin.Context, body interface{}) error
	//ValidValidation : 驗證表單資訊
	ValidValidation(v *validation.Validation) error
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
		return e.ParameterError("error parameter, please check your parameter again.")
	}
	return nil
}

//ValidValidation : 驗證表單資訊
func (r *RequestHandler) ValidValidation(v *validation.Validation) error {
	if v.HasErrors() {
		for _, err := range v.Errors {
			return e.ParameterError(fmt.Sprintf("parameter `%s` %s.", err.Key, err.Message))
		}
	}
	return nil
}

type RequesterInterface interface {
	GET(uri string) (*e.Response, error)
	POST(uri string, param interface{}) (*e.Response, error)
	PUT(uri string, param interface{}) (*e.Response, error)
	PATCH(uri string, param interface{}) (*e.Response, error)
	DELETE(uri string, param interface{}) (*e.Response, error)
	GetWithHeader(uri string, header map[string]string) (*e.Response, error)
	PostWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error)
	PutWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error)
	PatchWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error)
	DeleteWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error)
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

func (r *Requester) ReadJSONResponse(in *http.Response) (*e.Response, error) {
	defer in.Body.Close()
	var resp = new(e.Response)
	err := json.NewDecoder(in.Body).Decode(resp)
	return resp, err
}

func (r *Requester) GET(uri string) (*e.Response, error) {
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
func (r *Requester) POST(uri string, param interface{}) (*e.Response, error) {
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

func (r *Requester) PUT(uri string, param interface{}) (*e.Response, error) {
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

func (r *Requester) PATCH(uri string, param interface{}) (*e.Response, error) {
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

func (r *Requester) DELETE(uri string, param interface{}) (*e.Response, error) {
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

func (r *Requester) GetWithHeader(uri string, header map[string]string) (*e.Response, error) {
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
func (r *Requester) PostWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error) {
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
func (r *Requester) DeleteWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error) {
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
func (r *Requester) PutWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error) {
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
func (r *Requester) PatchWithHeader(uri string, header map[string]string, param interface{}) (*e.Response, error) {
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