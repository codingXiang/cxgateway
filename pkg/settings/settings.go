package settings

import (
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm/model"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"

	// "os"

)

//Struct
type (
	ConfigInterface interface {
		GetDatabase() model.DatabaseInterface
		GetRedis() model.RedisInterface
		GetApplication() ApplicationInterface
	}
	//Configuration : 整體設定檔
	Configuration struct {
		Database    *model.Database `yaml:"database"`    //資料庫相關參數
		Redis       *model.Redis    `yaml:"redis"`       //Redis相關參數
		Application *Application    `yaml:"application"` //應用程式相關參數
	}
)

var (
	//ConfigData : 設定檔變數
	ConfigData ConfigInterface
)

//NewConfigData 初始化
func NewConfigData() ConfigInterface {

	c := &Configuration{}
	switch os.Getenv("ACTIVE") {
	case "dev":
		c.GetConfig("config/config.test.yaml")
		break
	case "prod":
		c.GetConfig("config/config.yaml")
		break
	case "":
		log.Println("沒有設定環境變數 ACTIVE")
		break
	}
	return c
}

//GetConfig : 取得設定檔
func (c *Configuration) GetConfig(configPath string) error {
	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		logger.Log.Error("load yaml file error", err)
		return err
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln("transform yaml file error", err)
		return err
	}

	return nil
}

func (c *Configuration) GetDatabase() model.DatabaseInterface {
	return c.Database
}

func (c *Configuration) GetRedis() model.RedisInterface {
	return c.Redis
}

func (c *Configuration) GetApplication() ApplicationInterface {
	return c.Application
}
