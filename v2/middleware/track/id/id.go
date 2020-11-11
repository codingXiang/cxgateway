package id

import (
	"fmt"
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/codingXiang/cxgateway/v2/server"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"strings"
)

type RequestID struct {
	config *viper.Viper
}

func New(config *viper.Viper) middleware.Object {
	return &RequestID{
		config: config,
	}
}

func (c *RequestID) GetConfig() *viper.Viper {
	return c.config
}


func (r *RequestID) SetConfig(config *viper.Viper) {
	r.config = config
}

func (r *RequestID) Handle() gin.HandlerFunc {
	id := r.config.GetString(configer.GetConfigPath(server.Application, server.AppId))
	if id == "" {
		return func(c *gin.Context) {
			c.Next()
		}
	} else {
		return func(c *gin.Context) {
			trackID := uuid.NewV4()
			trackKey := fmt.Sprintf("X-%s-Track-Id", strings.ToUpper(id))
			c.Writer.Header().Set(trackKey, trackID.String())
			c.Next()
		}
	}
}
