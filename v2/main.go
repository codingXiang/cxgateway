package main

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware/auth"
	"github.com/codingXiang/cxgateway/v2/middleware/cache"
	"github.com/codingXiang/cxgateway/v2/middleware/i18n"
	"github.com/codingXiang/cxgateway/v2/middleware/logger"
	"github.com/codingXiang/cxgateway/v2/middleware/track/id"
	"github.com/codingXiang/cxgateway/v2/middleware/track/version"
	http2 "github.com/codingXiang/cxgateway/v2/module/auto_register/delivery/http"
	"github.com/codingXiang/cxgateway/v2/module/auto_register/repository"
	"github.com/codingXiang/cxgateway/v2/module/auto_register/service"
	"github.com/codingXiang/cxgateway/v2/server"
)

const (
	CONFIG_PATH = "./config"
	CONFIG      = "config"
	AUTH        = "auth"
	AUTO        = "registration"
)

func init() {
	configer.Config = configer.NewConfiger()
	config := configer.NewCore(configer.YAML, CONFIG, CONFIG_PATH)
	configer.Config.AddCore(CONFIG, config)

	auto := configer.NewCore(configer.YAML, AUTO, CONFIG_PATH)
	configer.Config.AddCore(AUTO, auto)

	auth := configer.NewCore(configer.YAML, AUTH, CONFIG_PATH)
	configer.Config.AddCore(AUTH, auth)
}

func main() {
	if config, err := configer.Config.GetCore(CONFIG).ReadConfig(); err == nil {
		server.Gateway = server.New(nil, config)
		server.Gateway.Use(
			log.New(config),
			version.New(config),
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
	server.Gateway.Run()
}
