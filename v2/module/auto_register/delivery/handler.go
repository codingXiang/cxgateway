package delivery

import (
	"github.com/gin-gonic/gin"
)

//HttpHandler http流量 handler
type HttpHandler interface {
	GetConfig(c *gin.Context) error
	Registered(c *gin.Context) error
	Initial(c *gin.Context) error
}

//GRPCHandler gRPC流量 handler
type GRPCHandler interface {
	//gRpcImplement
}

//CmdHandler cli handler
type CmdHandler interface {
	//CmdImplement
}
