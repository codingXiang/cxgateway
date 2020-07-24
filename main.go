package main

import (
	"github.com/codingXiang/configer"
	. "github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/cxgateway/module/version/delivery/http"
	"github.com/codingXiang/cxgateway/module/version/repository"
	"github.com/codingXiang/cxgateway/module/version/service"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/codingXiang/gogo-i18n"
)

func init() {
	//初始化 configer，設定預設讀取環境變數
	config := configer.NewConfigerCore("yaml", "config", "./config", ".")
	config.SetAutomaticEnv("")
	db := configer.NewConfigerCore("yaml", "storage", "./config", ".")
	db.SetAutomaticEnv("")
	redis := configer.NewConfigerCore("yaml", "storage", "./config", ".")
	redis.SetAutomaticEnv("")
	//初始化 Gateway
	Gateway = NewApiGateway("config", config)

	var err error

	//設定資料庫
	if orm.DatabaseORM, err = orm.NewOrm("database", db); err == nil {
		// 建立 Table Schema (Module)
		logger.Log.Debug("create table")
		{
			_ = orm.DatabaseORM.CheckTable(false, gogo_i18n.GoGoi18nMessage{})
			//工單模組
			{

			}
		}
	} else {
		logger.Log.Error(err.Error())
		panic(err.Error())
	}
	if orm.RedisORM, err = orm.NewRedisClient("redis", redis); err != nil {
		logger.Log.Error(err.Error())
		panic(err.Error())
	}
}

func main() {
	Gateway.EnableAutoRegistration("registration", "yaml", "./config")
	versionRepo := repository.NewVersionRepository(orm.DatabaseORM, orm.RedisORM)
	versionSvc := service.NewVersionService(versionRepo)
	http.NewVersionHttpHandler(Gateway, versionSvc)
	//運行 Gateway
	Gateway.Run()
}