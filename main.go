package main

import (
	"github.com/codingXiang/configer"
	. "github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/codingXiang/gogo-i18n"
)

func init() {
	//初始化 Gateway
	Gateway = NewApiGateway("test", nil)
	configer.Config.AddCore("storage", configer.NewConfigerCore("yaml", "storage", "./config", "."))
	if data, err := configer.Config.GetCore("storage").ReadConfig(); err == nil {
		//設定 Database 連線
		if setting := data.Get("database"); setting != nil {
			orm.NewOrm(orm.InterfaceToDatabase(setting))
			logger.Log.Debug("create table")
			{
				orm.DatabaseORM.CheckTable(false, gogo_i18n.GoGoi18nMessage{})
			}
		} else {
			logger.Log.Error("database setting is not exist")
			panic("must need to setting database config")

		}
		//設定 Redis 連線
		if setting := data.Get("redis"); setting != nil {
			orm.NewRedisClient(orm.InterfaceToRedis(setting))
		} else {
			logger.Log.Error("redis setting is not exist")
			panic("must need to setting redis config")
		}
	}
}

func main() {
	//運行 Gateway
	Gateway.Run()
}
