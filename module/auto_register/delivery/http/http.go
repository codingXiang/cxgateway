package http

import (
	"github.com/astaxie/beego/validation"
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/model"
	"github.com/codingXiang/cxgateway/module/auto_register"
	"github.com/codingXiang/cxgateway/module/auto_register/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/i18n"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	MODULE = "service_registered"
)

type AutoRegisteredHttpHandler struct {
	i18nMsg i18n.I18nMessageHandlerInterface
	gateway cx.HttpHandler
	svc     auto_register.Service
}

func NewAutoRegisteredHttpHandler(gateway cx.HttpHandler, svc auto_register.Service) delivery.HttpHandler {
	var handler = &AutoRegisteredHttpHandler{
		i18nMsg: i18n.NewI18nMessageHandler(MODULE),
		gateway: gateway,
		svc:     svc,
	}
	/*
		v1 版本的 AutoRegistered API
	*/
	api := gateway.GetApiRoute().Group("/register")
	api.GET("/config/:key", e.Wrapper(handler.GetConfig))
	api.GET("/initial", e.Wrapper(handler.Initial))
	api.POST("", e.Wrapper(handler.Registered))

	return handler
}

func (a *AutoRegisteredHttpHandler) GetConfig(c *gin.Context) error {
	a.i18nMsg.SetCore(util.GetI18nData(c))
	if data, err := a.svc.GetConfig(c.Params.ByName("key")); err == nil {
		c.JSON(a.i18nMsg.GetSuccess(data))
		return nil
	} else {
		return a.i18nMsg.GetError(err)
	}

}

func (g *AutoRegisteredHttpHandler) Registered(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  = new(model.ServiceRegister)
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetCore(util.GetI18nData(c))
	//綁定參數
	var err = c.ShouldBindWith(&data, binding.JSON)
	if err != nil || data == nil {
		return g.i18nMsg.ParameterFormatError()
	}

	//驗證表單資訊是否填寫充足
	valid.Required(&data.Name, "name")
	valid.Required(&data.URL, "url")

	if err := util.NewRequestHandler().ValidValidation(valid); err != nil {
		return err
	}

	if data, err := g.svc.Register(data); err == nil {
		c.JSON(g.i18nMsg.CreateSuccess(data))
		return nil
	} else {
		return g.i18nMsg.CreateError(err)
	}
}

func (g *AutoRegisteredHttpHandler) Initial(c *gin.Context) error {
	//將 middleware 傳入的 i18n 進行轉換
	g.i18nMsg.SetCore(util.GetI18nData(c))
	if err := g.svc.Initial(); err == nil {
		c.JSON(g.i18nMsg.GetSuccess(""))
		return nil
	} else {
		return g.i18nMsg.GetError(err)
	}
}
