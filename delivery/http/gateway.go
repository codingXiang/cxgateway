package http

import (
	"github.com/codingXiang/configer"
	"github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/middleware"
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

	if data, err := configer.Config.GetCore("config").ReadConfig(); err == nil {
		gateway.engine.Use(middleware.RequestIDMiddleware(data.GetString("application.appId"))).Use(middleware.Logger(), gin.Recovery())

	}


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
