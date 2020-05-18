package delivery

import (
	"github.com/codingXiang/configer"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	GetEngine() *gin.Engine
	GetApiRoute() *gin.RouterGroup
	GetHandler() util.RequestHandlerInterface
	GetConfig() configer.CoreInterface
	GetUploadPath() string
	Run()
}
