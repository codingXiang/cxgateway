package version

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/codingXiang/cxgateway/v2/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type RequestVersion struct {
	config *viper.Viper
}

func New(config *viper.Viper) middleware.Object {
	return &RequestVersion{
		config: config,
	}
}

//Version
func (r *RequestVersion) Handle() gin.HandlerFunc {
	enable := r.config.GetBool(configer.GetConfigPath(server.Application, server.Version, server.Enable))
	if enable {
		key := r.config.GetString(configer.GetConfigPath(server.Application, server.Version, server.Key))
		value := r.config.GetString(configer.GetConfigPath(server.Application, server.Version, server.Value))
		return func(c *gin.Context) {
			c.Writer.Header().Set(key, value)
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			c.Next()
		}
	}
}
