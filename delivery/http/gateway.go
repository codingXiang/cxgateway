package http

import (
	"fmt"
	"github.com/codingXiang/configer"
	"github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/middleware"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/gogo-i18n"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApiGateway struct {
	engine     *gin.Engine
	Api        *gin.RouterGroup
	handler    util.RequestHandlerInterface
	configName string
}

var (
	Gateway delivery.HttpHandler
)

func NewApiGateway(configName string, core configer.CoreInterface) delivery.HttpHandler {
	var (
		//config  = settings.ConfigData.Data.Application
		gateway = &ApiGateway{
			engine: gin.Default(),
		}
	)
	if configName != "" {
		gateway.configName = configName
	} else {
		gateway.configName = "default"
	}
	//初始化 configer
	configer.Config = configer.NewConfiger()
	//設定多語系 Handler
	gogo_i18n.LangHandler = gogo_i18n.NewLanguageHandler()
	//設定預設資料
	if core == nil{
		if gateway.configName == "default" {
			configer.Config.AddCore(gateway.configName, configer.NewConfigerCore("yaml", "config", "./config", "."))
			configer.Config.GetCore(gateway.configName).SetAutomaticEnv()
			configer.Config.GetCore(gateway.configName).SetDefault("i18n", map[string]interface{}{
				"defaultLanguage": "zh_Hant",
				"file": map[string]string{
					"path": "./i18n",
					"type": "yaml",
				},
			})
			configer.Config.GetCore(gateway.configName).
				SetDefault("application", map[string]interface{}{
					"timeout": map[string]int{
						"read":  1000,
						"write": 1000,
					},
					"port":  8080,
					"appId": "app",
					"mode":  "debug",
					"log": map[string]string{
						"level":  "debug",
						"format": "json",
					},
					"appToken":     "defaultToken",
					"appBaseRoute": "/api",
				})
		}
	} else {
		configer.Config.AddCore(gateway.configName, configer.NewConfigerCore("yaml", "config", "./config", "."))
		configer.Config.GetCore(gateway.configName).SetAutomaticEnv()
	}

	gateway.handler = util.NewRequestHandler()

	if data, err := gateway.GetConfig().ReadConfig(); err == nil {
		//設定 log 等級與格式
		logger.Log = logger.NewLogger(logger.InterfaceToLogger(data.Get("application.log")))
		gateway.engine.
			Use(middleware.Logger(), gin.Recovery()).
			Use(middleware.RequestIDMiddleware(data.GetString("application.appId"))).
			Use(middleware.GoI18nMiddleware(data))
		gateway.Api = gateway.engine.Group(data.GetString("application.apiBaseRoute"))
	} else {
		panic(fmt.Sprintf("config %s is not set", gateway.configName))
	}

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

func (this *ApiGateway) GetConfig() configer.CoreInterface {
	return configer.Config.GetCore(this.configName)
}

func (this *ApiGateway) Run() {
	if data, err := this.GetConfig().ReadConfig(); err == nil {
		var (
			port         = data.GetInt("application.port")          //伺服器的 port
			writeTimeout = data.GetInt("application.timeout.write") //伺服器的寫入超時時間
			readTimeout  = data.GetInt("application.timeout.read")  //伺服器讀取超時時間
			mode         = data.GetString("application.mode")             //伺服器模式
		)
		//設定運行模式
		if mode == "release" {
			gin.SetMode(mode)
		}
		logger.Log.Info("Setting Http Server Info")
		logger.Log.Debug("server port = ", port)
		// 設定 http server
		server := &http.Server{
			Addr:           fmt.Sprintf(":%d", port),
			Handler:        Gateway.GetEngine(),
			ReadTimeout:    time.Duration(readTimeout) * time.Second,
			WriteTimeout:   time.Duration(writeTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.Log.Info("API Gateway Start Running")
		//啟動 http server
		server.ListenAndServe()
	}
}
