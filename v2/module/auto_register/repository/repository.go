package repository

import (
	"encoding/json"
	"github.com/codingXiang/cxgateway/v2/model"
	"github.com/codingXiang/cxgateway/v2/module/auto_register"
	"github.com/codingXiang/cxgateway/v2/util/util"
	"github.com/codingXiang/go-logger/v2"
	"github.com/codingXiang/go-orm/v2/redis"
	"github.com/spf13/viper"
	"strings"
)

var AutoRegisteredRedisClient *redis.RedisClient

type AutoRegisteredRepository struct {
	data   *viper.Viper
	Client *redis.RedisClient
}

func NewAutoRegisteredRepository(config *viper.Viper) (auto_register.Repository, error) {
	client := redis.New(config)
	if _, err := client.GetInfo(); err != nil{
		logger.Log.Error("connect to auto registration redis failed, err =", err.Error())
		return nil, err
	} else {
		AutoRegisteredRedisClient = client
	}
	return &AutoRegisteredRepository{
		Client: AutoRegisteredRedisClient,
		data:   config,
	}, nil
}

func (a *AutoRegisteredRepository) GetConfig(key string) (string, error) {
	return a.Client.GetValue(key)
}

func (a *AutoRegisteredRepository) Register(data *model.ServiceRegister) (*model.ServiceRegister, error) {
	err := a.Client.SetKeyValue(data.Name, data.URL, 0)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AutoRegisteredRepository) toAutoRegistrationInfo(data interface{}) (*model.AutoRegistrationInfo, error) {
	var result *model.AutoRegistrationInfo
	tmp, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tmp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AutoRegisteredRepository) Initial() error {
	logger.Log.Info("start auto service registration")
	requester := util.NewRequester(nil)
	registeredPath := a.data.GetString("registeredPath")

	if a.data.GetBool("local.startInit") {
		//local
		localObj := &model.ServiceRegister{a.data.GetString("local.name"), a.data.GetString("local.url")}
		for _, destination := range strings.Split(a.data.GetString("local.destinations"), ",") {
			url := destination + registeredPath
			logger.Log.Info("register", url, "name =", localObj.Name, "url =", localObj.URL)
			_, err := requester.POST(url, localObj)
			if err != nil {
				logger.Log.Error("auto service registration local failed, err =", err.Error())
				return err
			}
		}
	} else {
		logger.Log.Info("not auto registered local")
	}

	if a.data.GetBool("remote.startInit") {
		//remote
		remoteObj := &model.ServiceRegister{a.data.GetString("remote.name"), a.data.GetString("remote.url")}
		for _, destination := range strings.Split(a.data.GetString("remote.destinations"), ",") {
			url := destination + registeredPath
			logger.Log.Info("register", url, "name =", remoteObj.Name, "url =", remoteObj.URL)
			_, err := requester.POST(url, remoteObj)
			if err != nil {
				logger.Log.Error("auto service registration remote failed, err =", err.Error())
				return err
			}
		}
	} else {
		logger.Log.Info("not auto registered remote")
	}

	logger.Log.Info("auto service registration success")
	return nil
}
