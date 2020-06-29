# 封裝 GORM 的 ORM
## 如何使用
### 載入模組
```shell
go get -u github.com/codingXiang/go-orm
```

### 設定 ORM 實例
#### Database
```
var err error
//初始化 logger
logger.Log = logger.NewLogger(logger.Logger{
	Level:  "debug",
	Format: "json",
})
//設定 configer
databaseConfig := configer.NewConfigerCore("yaml", "config", "./example")
//建立 orm instance
if orm.DatabaseORM, err = orm.NewOrm("database", databaseConfig); err != nil {
	panic(err)
}
//取得實例
orm.DatabaseORM.GetInstance()
//版本更新
if err = orm.DatabaseORM.Upgrade(&Test{}); err != nil {
	panic(err.Error())
}
```

#### Redis
```
/*
	設定 Logger
 */
logger.Log = logger.NewLogger(logger.Logger{
	Level:  "debug",
	Format: "json",
})
//設定 configer
config := configer.NewConfigerCore("yaml", "redis-config", "./example")

/*
	建立實例
 */
var err error
if orm.RedisORM, err = orm.NewRedisClient("redis", config); err != nil {
	panic(err.Error())
}
//上傳 key
orm.RedisORM.SetKeyValue("test", "test", 0)
```
## 參數設定
可以參考 example 裡面的 config.yaml，此格式可對照 model 裡面的 database
