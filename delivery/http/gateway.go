package http

import (
	"github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/middleware"
	"github.com/codingXiang/cxgateway/pkg/settings"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/gin-gonic/gin"
)

type ApiGateway struct {
	engine  *gin.Engine
	Api     *gin.RouterGroup
	handler util.RequestHandlerInterface
}

func NewApiGateway() delivery.HttpHandler {
	var (
		//config  = settings.ConfigData.Data.Application
		gateway = &ApiGateway{
			engine: gin.Default(),
		}
	)
	gateway.handler = util.NewRequestHandler()
	gateway.engine.Use(middleware.RequestIDMiddleware(settings.ConfigData.GetApplication().GetAppID())).Use(middleware.Logger(), gin.Recovery())

	gateway.Api = gateway.engine.Group("/api")
	return gateway
}

func (gateway *ApiGateway) GetEngine() *gin.Engine {
	return gateway.engine
}

func (gateway *ApiGateway) GetHandler() util.RequestHandlerInterface {
	return gateway.handler
}

func (gateway *ApiGateway) GetApiRoute() *gin.RouterGroup {
	return gateway.Api
}
