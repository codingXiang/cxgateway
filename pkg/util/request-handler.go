package util

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/gin-gonic/gin"
)

type RequestHandlerInterface interface {
	//BindBody : 綁定 body
	BindBody(c *gin.Context, body interface{}) error
	//ValidFormField : 驗證表單資訊
	ValidFormField(data interface{}) error
}

type RequestHandler struct {
	context *gin.Context
}

func NewRequestHandler() RequestHandlerInterface {
	return &RequestHandler{}
}

func (r *RequestHandler) BindBody(c *gin.Context, body interface{}) error {
	var err = c.Bind(&body)
	if err != nil {
		return e.ParameterError("error parameter, please check your parameter again.")
	}
	return nil
}

func (r *RequestHandler) ValidFormField(data interface{}) error {
	var v = &validation.Validation{}
	var hasError, err = v.Valid(&data)
	if err != nil {
		return e.UnknownError(fmt.Sprintf("valid form something error = %s", err.Error()))
	}
	if hasError {
		for _, err := range v.Errors {
			return e.ParameterError(fmt.Sprintf("parameter `%s` %s.", err.Key, err.Message))
		}
	}
	return nil
}


func (r *RequestHandler) ValidValidation(v *validation.Validation) error {
	if v.HasErrors() {
		for _, err := range v.Errors {
			return e.ParameterError(fmt.Sprintf("parameter `%s` %s.", err.Key, err.Message))
		}
	}
	return nil
}