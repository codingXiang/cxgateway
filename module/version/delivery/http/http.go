package http

import (
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/module/version"
	"github.com/codingXiang/cxgateway/module/version/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/i18n"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	MODULE = "version"
)

type VersionHttpHandler struct {
	i18nMsg i18n.I18nMessageHandlerInterface
	gateway cx.HttpHandler
	svc     version.Service
}

func NewVersionHttpHandler(gateway cx.HttpHandler, svc version.Service) delivery.HttpHandler {
	var handler = &VersionHttpHandler{
		i18nMsg: i18n.NewI18nMessageHandler(MODULE),
		gateway: gateway,
		svc:     svc,
	}
	/*
		v1 版本的 Ticket API
	*/
	v1 := gateway.GetEngine().Group("")
	v1.GET("", e.Wrapper(handler.GetServerVersion))
	//v1.GET("/check", e.Wrapper(handler.CheckVersion))
	//v1.POST("/upgrade", e.Wrapper(handler.Upgrade))
	return handler
}

func (this VersionHttpHandler) GetServerVersion(c *gin.Context) error {
	this.i18nMsg.SetCore(util.GetI18nData(c))
	if version, err := this.svc.GetServerVersion(); err != nil {
		return this.i18nMsg.GetError(err)
	} else {
		c.JSON(this.i18nMsg.GetSuccess(version))
	}
	return nil
}

func (this VersionHttpHandler) CheckVersion(c *gin.Context) error {
	this.i18nMsg.SetCore(util.GetI18nData(c))
	if err := this.svc.CheckVersion(); err != nil {
		return this.i18nMsg.GetError(err)
	} else {
		c.JSON(this.i18nMsg.GetSuccess(nil))
	}
	return nil
}

func (this VersionHttpHandler) Upgrade(c *gin.Context) error {
	this.i18nMsg.SetCore(util.GetI18nData(c))
	if err := this.svc.Upgrade(); err != nil {
		return this.i18nMsg.UpdateError(err)
	} else {
		c.JSON(this.i18nMsg.UpdateSuccess(nil))
	}
	return nil
}
