package main

import (
	"fmt"
	. "github.com/codingXiang/cxgateway/delivery/http"
	. "github.com/codingXiang/cxgateway/pkg/settings"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	//取得組態變數
	ConfigData = NewConfigData()
	//設定 log 等級與格式
	logger.Log = logger.NewLogger(ConfigData.GetApplication().GetLog())
	//設定 Database 連線
	orm.NewOrm(ConfigData.GetDatabase())
	//設定 Redis 連線
	orm.NewRedisClient(ConfigData.GetRedis())
	//設定運行模式
	mode := ConfigData.GetApplication().GetMode()
	// port := settings.ConfigData.Data.Application.Port
	if mode == "release" {
		gin.SetMode("release")
	}
}

func main() {
	// General Variable
	var (
		timeout      = ConfigData.GetApplication().GetTimeout()
		readTimeout  = timeout.GetRead()
		writeTimeout = timeout.GetWrite()
	)
	// 建立 API Gateway
	logger.Log.Debug("Create API Gateway")
	var (
		gateway = NewApiGateway()
	)
	logger.Log.Info("Setting Http Server Info")
	// 設定 http server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", ConfigData.GetApplication().GetPort()),
		Handler:        gateway.GetEngine(),
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Log.Info("API Gateway Start Running")
	//啟動 http server
	s.ListenAndServe()
}
