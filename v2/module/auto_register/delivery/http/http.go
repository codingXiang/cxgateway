package http

import (
	"github.com/astaxie/beego/validation"
	"github.com/codingXiang/cxgateway/v2/model"
	"github.com/codingXiang/cxgateway/v2/module/auto_register"
	"github.com/codingXiang/cxgateway/v2/server"
	"github.com/codingXiang/cxgateway/v2/util/e"
	"github.com/codingXiang/cxgateway/v2/util/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	MODULE = "service_registered"
)

type AutoRegisteredHttpHandler struct {
	server.HttpModule
	svc auto_register.Service
}

func NewAutoRegisteredHttpHandler(svc auto_register.Service) *AutoRegisteredHttpHandler {
	var handler = &AutoRegisteredHttpHandler{
		HttpModule: new(server.Http),
		svc:        svc,
	}
	handler.HttpModule.SetI18n(MODULE)
	/*
		v1 版本的 AutoRegistered API
	*/

	return handler
}

func (g *AutoRegisteredHttpHandler) Setup() {
	api := g.GetGateway().GetApiRoute().Group("/register")
	api.GET("/config/:key", e.Wrapper(g.GetConfig))
	api.GET("/initial", e.Wrapper(g.Initial))
	api.POST("", e.Wrapper(g.Registered))
}

func (a *AutoRegisteredHttpHandler) GetConfig(c *gin.Context) error {
	a.HttpModule.GetI18n().SetCore(util.GetI18nData(c))
	if data, err := a.svc.GetConfig(c.Params.ByName("key")); err == nil {
		c.JSON(a.GetI18n().GetSuccess(data))
		return nil
	} else {
		return a.GetI18n().GetError(err)
	}
}

func (g *AutoRegisteredHttpHandler) Registered(c *gin.Context) error {
	var (
		valid = new(validation.Validation)
		data  = new(model.ServiceRegister)
	)
	//將 middleware 傳入的 i18n 進行轉換
	g.GetI18n().SetCore(util.GetI18nData(c))
	//綁定參數
	var err = c.ShouldBindWith(&data, binding.JSON)
	if err != nil || data == nil {
		return g.GetI18n().ParameterFormatError()
	}

	//驗證表單資訊是否填寫充足
	valid.Required(&data.Name, "name")
	valid.Required(&data.URL, "url")

	if err := util.NewRequestHandler().ValidValidation(valid); err != nil {
		return err
	}

	if data, err := g.svc.Register(data); err == nil {
		c.JSON(g.GetI18n().CreateSuccess(data))
		return nil
	} else {
		return g.GetI18n().CreateError(err)
	}
}

func (g *AutoRegisteredHttpHandler) Initial(c *gin.Context) error {
	//將 middleware 傳入的 i18n 進行轉換
	g.GetI18n().SetCore(util.GetI18nData(c))
	if err := g.svc.Initial(); err == nil {
		c.JSON(g.GetI18n().GetSuccess(""))
		return nil
	} else {
		return g.GetI18n().GetError(err)
	}
}
