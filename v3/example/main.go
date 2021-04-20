package main

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v3/middleware/cache"
	"github.com/codingXiang/cxgateway/v3/middleware/cors"
	"github.com/codingXiang/cxgateway/v3/middleware/logger"
	service "github.com/codingXiang/cxgateway/v3/middleware/module"
	"github.com/codingXiang/cxgateway/v3/middleware/pagination"
	"github.com/codingXiang/cxgateway/v3/middleware/track/id"
	"github.com/codingXiang/cxgateway/v3/middleware/track/version"
	"github.com/codingXiang/cxgateway/v3/server"
	response "github.com/codingXiang/cxgateway/v3/util/response"
	"github.com/codingXiang/go-logger"
	"github.com/gin-gonic/gin"
)

const (
	ConfigPath = "./example/config"
	Config     = "config"
)

func init() {
	configer.Config = configer.NewConfiger()
	config := configer.NewCore(configer.YAML, Config, ConfigPath)
	configer.Config.AddCore(Config, config)
}

const appId response.ServiceCode = "05"

func main() {
	if config, err := configer.Config.GetCore(Config).ReadConfig(); err == nil {
		response.SetCurrentServiceCode(appId)
		logger.Log = logger.NewLoggerWithConfiger(config)
		server.Gateway = server.New(nil, config)
		server.Gateway.Use(
			cors.New(nil),
			log.New(nil),
			version.New(nil),
			id.New(config),
			cache.New(),
			pagination.New(),
		)
	} else {
		panic(err.Error())
	}
	test := server.Gateway.GetEngine().Group("/api/v1/test")
	var testModule response.ModuleCode = "01"
	test.Use(service.New(testModule).Handle())
	test.GET("", func(c *gin.Context) {
		c.Set(response.MessageKey, "test ok")
		c.Set(response.DataKey, "hello")
		c.JSON(response.StatusOK(c))
	})
	test.POST("", func(c *gin.Context) {
		c.Set(response.MessageKey, "test ok")
		c.Set(response.DataKey, "hello")
		c.JSON(response.StatusOK(c))
	})
	server.Gateway.Run()
}

