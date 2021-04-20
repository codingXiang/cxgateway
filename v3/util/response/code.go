package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Code int

type Status string

const (
	Failed  Status = "1"
	Success        = "2"
)

type ServiceCode string

const (
	UserService    ServiceCode = "01"
	ProductService ServiceCode = "02"
	OrderService   ServiceCode = "03"
)

var currentService ServiceCode

func SetCurrentServiceCode(code ServiceCode) {
	currentService = code
}

type ModuleCode string

const (
	notRegistry ModuleCode = "99"
)

type MethodCode string

const (
	Get    MethodCode = "0"
	Post              = "1"
	Put               = "2"
	Patch             = "3"
	Delete            = "4"
)

func GetMethodCode(in string) MethodCode {
	switch in {
	case "GET":
		return Get
	case "POST":
		return Post
	case "PUT":
		return Put
	case "PATCH":
		return Patch
	case "DELETE":
		return Delete
	default:
		return Get
	}
}

type CustomCode string

func GetCode(status Status, service ServiceCode, module ModuleCode, method MethodCode, customCode ...string) Code {
	code := fmt.Sprintf("%s%s%s%s", status, service, module, method)
	for _, c := range customCode {
		code += c
	}
	res, _ := strconv.Atoi(code)
	return (Code)(res)
}

func NewCode(c *gin.Context) Code {
	var (
		codeStatus  Status     = Success
		moduleCode  ModuleCode = "00"
		serviceCode            = currentService
	)
	if in, exist := c.Get(CodeStatus); exist {
		codeStatus = in.(Status)
	}
	if in, exist := c.Get(ModuleName); exist {
		moduleCode = in.(ModuleCode)
	}
	return GetCode(codeStatus, serviceCode, moduleCode, GetMethodCode(c.Request.Method))
}
