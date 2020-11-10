package delivery

import (
	"github.com/codingXiang/configer"
	"github.com/codingXiang/cxgateway/v2/pkg/util"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	GetEngine() *gin.Engine
	GetApiRoute() *gin.RouterGroup
	GetHandler() util.RequestHandlerInterface
	GetConfig() configer.CoreInterface
	GetUploadPath() string
	EnableAutoRegistration(configName string, configType string, configPath ...string) error
	Run()
}
