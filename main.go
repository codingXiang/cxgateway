package main

import (
	"fmt"
	"github.com/codingXiang/configer"
	. "github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	//取得組態變數
	configer.Config = configer.NewConfiger().
		AddCore("config", configer.NewConfigerCore("yaml", "config", "./config", ".")).
		AddCore("cloud", configer.NewConfigerCore("json", "cloud", "./config", ".")).
		AddCore("java", configer.NewConfigerCore("properties", "java", "./config", "."))

	if data, err := configer.Config.GetCore("config").ReadConfig(); err == nil {
		//設定 log 等級與格式
		logger.Log = logger.NewLogger(logger.InterfaceToLogger(data.Get("application.log")))
		//設定 Database 連線
		orm.NewOrm(orm.InterfaceToDatabase(data.Get("database")))
		//設定 Redis 連線
		orm.NewRedisClient(orm.InterfaceToRedis(data.Get("redis")))
		//設定運行模式
		mode := data.Get("application.mode")
		// port := settings.ConfigData.Data.Application.Port
		if mode == "release" {
			gin.SetMode("release")
		}
	}
}

func main() {
	if data, err := configer.Config.GetCore("config").ReadConfig(); err == nil {
		// 建立 API Gateway
		logger.Log.Debug("Create API Gateway")
		var (
			gateway = NewApiGateway()
		)
		logger.Log.Info("Setting Http Server Info")
		// 設定 http server
		s := &http.Server{
			Addr:           fmt.Sprintf(":%d", data.GetInt("application.port")),
			Handler:        gateway.GetEngine(),
			ReadTimeout:    time.Duration(data.GetInt("application.timeout.write")) * time.Second,
			WriteTimeout:   time.Duration(data.GetInt("application.timeout.read")) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.Log.Info("API Gateway Start Running")
		//啟動 http server
		s.ListenAndServe()
	}

}
