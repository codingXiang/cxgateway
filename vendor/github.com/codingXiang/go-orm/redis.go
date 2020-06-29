package orm

import (
	"encoding/json"
	"fmt"
	"github.com/codingXiang/configer"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm/model"
	"github.com/go-redis/redis"
	"time"
)

type (
	//RedisClientInterface : Redis客戶端介面
	RedisClientInterface interface {
		GetInfo() (string, error)
		SetKeyValue(key string, value interface{}, expiration time.Duration) error
		GetValue(key string) (string, error)
		RemoveKey(key string) error
	}
	//RedisClient : Redis客戶端
	RedisClient struct {
		client     *redis.Client
		prefix     string
		configName string
	}
)

var (
	RedisORM RedisClientInterface
)

func InterfaceToRedis(data interface{}) model.RedisInterface {
	var result = &model.Redis{}
	if jsonStr, err := json.Marshal(data); err == nil {
		json.Unmarshal(jsonStr, &result)
	}
	return result
}

//NewRedisClient : 建立 Redis Client 實例
func NewRedisClient(configName string, core configer.CoreInterface) (*RedisClient, error) {
	var (
		rc = &RedisClient{
			configName: configName,
		}
	)

	if configer.Config == nil {
		//初始化 configer
		configer.Config = configer.NewConfiger()
	}

	//加入 config
	configer.Config.AddCore(rc.configName, core)
	//讀取 config
	if data, err := configer.Config.GetCore(rc.configName).ReadConfig(nil); err == nil {
		var (
			url      = data.GetString("redis.url")
			port     = data.GetInt("redis.port")
			password = data.GetString("redis.password")
			db       = data.GetInt("redis.db")
			prefix   = data.GetString("redis.prefix")
		)
		//設定連線資訊
		option := &redis.Options{
			Addr: fmt.Sprintf("%s:%d", url, port),
			DB:   db,
		}
		rc.prefix = prefix
		if password != "" {
			option.Password = password
		}
		rc.client = redis.NewClient(option)
		logger.Log.Debug("check redis ...", rc.client)
		_, err = rc.GetInfo()
		if err != nil {
			errMsg := "redis connect error"
			logger.Log.Error(errMsg, err)
			return nil, err
		} else {
			logger.Log.Info("redis connect success")
			return rc, nil
		}
	} else {
		return nil, err
	}
}

//GetRedisInfo 取得 Redis 資訊
func (r *RedisClient) GetInfo() (string, error) {
	return r.client.Ping().Result()
}

//SetKeyValueWithExpire : 設定 Key 與 Value
func (r *RedisClient) SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(r.prefix+key, value, expiration).Err()
	return err
}

//GetValue : 取得 Key 的 Value
func (r *RedisClient) GetValue(key string) (string, error) {
	val := r.client.Get(r.prefix+key)
	return val.Val(), val.Err()
}

//RemoveKey : 刪除 Key
func (r *RedisClient) RemoveKey(key string) error {
	return r.client.Del(r.prefix+key).Err()
}
