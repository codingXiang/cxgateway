package main

import (
	"github.com/codingXiang/configer/v2"
	example2 "github.com/codingXiang/cxgateway/v2/example/grpc/pb"
	grpc2 "github.com/codingXiang/cxgateway/v2/example/grpc/server"
	"github.com/codingXiang/cxgateway/v2/middleware/auth"
	"github.com/codingXiang/cxgateway/v2/middleware/cache"
	"github.com/codingXiang/cxgateway/v2/middleware/cors"
	"github.com/codingXiang/cxgateway/v2/middleware/i18n"
	"github.com/codingXiang/cxgateway/v2/middleware/logger"
	"github.com/codingXiang/cxgateway/v2/middleware/track/id"
	"github.com/codingXiang/cxgateway/v2/middleware/track/version"
	http2 "github.com/codingXiang/cxgateway/v2/module/auto_register/delivery/http"
	"github.com/codingXiang/cxgateway/v2/module/auto_register/repository"
	"github.com/codingXiang/cxgateway/v2/module/auto_register/service"
	"github.com/codingXiang/cxgateway/v2/module/service_discovery"
	"github.com/codingXiang/cxgateway/v2/server"
	"github.com/codingXiang/cxgateway/v2/server/grpc"
	"github.com/codingXiang/go-logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CONFIG_PATH       = "./config"
	CONFIG            = "config"
	AUTH              = "auth"
	AUTO              = "registration"
	GRPC              = "grpc"
	SERVICE_DISCOVERY = "service_discovery"
)

func init() {
	configer.Config = configer.NewConfiger()
	config := configer.NewCore(configer.YAML, CONFIG, CONFIG_PATH)
	configer.Config.AddCore(CONFIG, config)

	auto := configer.NewCore(configer.YAML, AUTO, CONFIG_PATH)
	configer.Config.AddCore(AUTO, auto)

	auth := configer.NewCore(configer.YAML, AUTH, CONFIG_PATH)
	configer.Config.AddCore(AUTH, auth)

	grpc := configer.NewCore(configer.YAML, GRPC, CONFIG_PATH)
	configer.Config.AddCore(GRPC, grpc)
}

func main() {

	service_discovery.Init(configer.YAML, SERVICE_DISCOVERY, CONFIG_PATH)
	service_discovery.StartWatch("/service/backend/")
	if config, err := configer.Config.GetCore(CONFIG).ReadConfig(); err == nil {
		logger.Log = logger.NewLoggerWithConfiger(config)
		server.Gateway = server.New(nil, config)
		server.Gateway.Use(
			auth.New(server.Gateway.GetAppID(), nil),
			cors.New(nil),
			log.New(nil),
			version.New(nil),
			id.New(config),
			i18n.New(config),
			cache.New(),
		)
	} else {
		panic(err.Error())
	}

	// 驗證
	if config, err := configer.Config.GetCore(AUTH).ReadConfig(); err == nil {
		server.Gateway.Use(auth.New(server.Gateway.GetAppID(), config))
	} else {
		panic(err.Error())
	}
	if config, err := configer.Config.GetCore(AUTO).ReadConfig(); err == nil {
		repo, err := repository.NewAutoRegisteredRepository(config)
		if err != nil {
			panic(err.Error())

		}
		repo.Initial()

		svc := service.NewAutoRegisteredService(repo)
		// 加入模組
		server.Gateway.AddModule(
			http2.NewAutoRegisteredHttpHandler(svc),
		)

	} else {
		panic(err.Error())
	}

	// grpc service 範例
	{
		if config, err := configer.Config.GetCore(GRPC).ReadConfig(); err == nil {
			grpc.Gateway = grpc.New(nil, config)
		}
		example := grpc2.NewExampleService()
		example2.RegisterExampleServiceServer(grpc.Gateway.GetServer(), example)
		grpc.Gateway.RunBackground()
	}
	server.Gateway.GetEngine().Use(func(c *gin.Context) {
		method := c.Request.Method
		url := c.Request.URL
		logger.Log.Debug(method, url)
	}).GET("/api/v1/test/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "test")
	})
	server.Gateway.Run()
}
