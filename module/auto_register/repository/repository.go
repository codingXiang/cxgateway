package repository

import (
	"encoding/json"
	"github.com/codingXiang/configer"
	"github.com/codingXiang/cxgateway/model"
	"github.com/codingXiang/cxgateway/module/auto_register"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type AutoRegisteredRepository struct {
	data   *viper.Viper
	client orm.RedisClientInterface
}

func NewAutoRegisteredRepository(config configer.CoreInterface) (auto_register.Repository, error) {
	client, err := orm.NewRedisClient("auto_registration", config)
	if err != nil {
		logger.Log.Error("connect to auto registration redis failed, err =", err.Error())
		return nil, err
	}
	if data, err := config.ReadConfig(nil); err == nil {
		return &AutoRegisteredRepository{
			client: client,
			data:   data,
		}, nil
	} else {
		return nil, err
	}
}

func (a *AutoRegisteredRepository) GetConfig(key string) (*model.ServiceRegister, error) {
	var result *model.ServiceRegister
	val, err := a.client.GetValue(key)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(val), &result); err == nil {
		return result, nil
	} else {
		return nil, err
	}
}

func (a *AutoRegisteredRepository) Register(data *model.ServiceRegister) (*model.ServiceRegister, error) {
	in, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = a.client.SetKeyValue(data.Name, string(in), 0)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *AutoRegisteredRepository) Initial() error {
	logger.Log.Info("start auto service registration")
	requester := util.NewRequester(nil)
	registeredPath := a.data.GetString("registeredPath")
	var infos = []*model.AutoRegistrationInfo{}

	data := a.data.Get("auto-registered")
	tmp, err := yaml.Marshal(data)
	if err != nil {
		logger.Log.Error("auto service registration init failed, err =", err.Error())
		return err
	}
	err = yaml.Unmarshal(tmp, &infos)
	if err != nil {
		logger.Log.Error("auto service registration init failed, err =", err.Error())
		return err
	}
	for _, info := range infos {
		obj := &model.ServiceRegister{info.Name, info.Url}
		for _, destination := range info.Destinations {
			url := destination + registeredPath
			_, err := requester.POST(url, obj)
			if err != nil {
				logger.Log.Error("auto service registration failed, err =", err.Error())
				return err
			}
		}
	}
	logger.Log.Info("auto service registration success")
	return nil
}
