package cors

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v3/middleware"
	"github.com/codingXiang/cxgateway/v3/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

type Cors struct {
	config *viper.Viper
}

func New(config *viper.Viper) middleware.Object {
	return &Cors{
		config: config,
	}

}

func (c *Cors) GetConfig() *viper.Viper {
	return c.config
}


func (c *Cors) SetConfig(config *viper.Viper) {
	c.config = config
}

func (c *Cors) Handle() gin.HandlerFunc {
	if c.config.GetBool(server.Default_) {
		return cors.Default()
	} else {
		return cors.New(NewConfig(c.config))
	}
}

func NewConfig(data *viper.Viper) cors.Config {
	if data.GetStringMap(server.Cors) == nil {
		return cors.DefaultConfig()
	}
	allowAllOrigin := data.GetBool(configer.GetConfigPath(server.Cors, server.AllowAllOrigin))
	config := cors.DefaultConfig()
	if allowAllOrigin {
		config.AllowAllOrigins = allowAllOrigin
	} else {
		config.AllowOrigins = strings.Split(data.GetString(configer.GetConfigPath(server.Cors, server.AllowOrigins)), ",")
	}
	config.AllowHeaders = strings.Split(data.GetString(configer.GetConfigPath(server.Cors, server.AllowHeaders)), ",")
	config.AllowMethods = strings.Split(data.GetString(configer.GetConfigPath(server.Cors, server.AllowMethods)), ",")
	return config
}
