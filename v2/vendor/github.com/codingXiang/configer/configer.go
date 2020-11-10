package configer

import (
	"bytes"
	"github.com/spf13/viper"
	"strings"
)

type (
	ConfigerInterface interface {
		GetCore(key string) CoreInterface
		AddCore(key string, handler CoreInterface) ConfigerInterface
	}
	CoreInterface interface {
		SetAutomaticEnv(prefix string)
		SetDefault(key string, value interface{}) CoreInterface
		WriteConfig() error
		WriteConfigAs(path string) error
		SetConfigType(in string)
		SetConfigName(in string)
		AddConfigPath(in string)
		ReadConfig(data []byte) (*viper.Viper, error)
	}

	//Configer : 整體設定檔
	Configer struct {
		handler map[string]CoreInterface
	}

	Core struct {
		core *viper.Viper
	}
)

var (
	//Config : 設定檔變數
	Config ConfigerInterface
)

// 參數依序為：
/// 1. 設定檔類型 (支援 yaml、yml、json、properties、ini、hcl、toml)
/// 2. 檔案名稱 (例如檔名為 config.yaml 就輸入 config)
/// 3. 後續皆為檔案路徑，可以支援多個路徑尋找檔案
func NewConfigerCore(configType string, configName string, paths ...string) CoreInterface {
	var handler = &Core{
		core: viper.New(),
	}
	handler.SetConfigType(configType)
	if configName != "" {
		handler.SetConfigName(configName)
		for _, path := range paths {
			handler.AddConfigPath(path)
		}
	}

	return handler
}

//NewConfiger 初始化
func NewConfiger() ConfigerInterface {
	var config = &Configer{}
	config.handler = map[string]CoreInterface{}
	return config
}

//GetCore : 取得組態控制器
func (this *Configer) GetCore(key string) CoreInterface {
	// 檢查 key 是否存在
	if value, ok := this.handler[key]; ok {
		return value
	}
	return nil
}

//AddCore : 加入組態控制器
func (this *Configer) AddCore(key string, handler CoreInterface) ConfigerInterface {
	this.handler[key] = handler
	return this
}
func (this *Core) SetAutomaticEnv(prefix string) {
	if prefix != "" {
		this.getCore().SetEnvPrefix(prefix)
	}
	this.getCore().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	this.getCore().AutomaticEnv()
}

func (this *Core) SetConfigType(in string) {
	this.getCore().SetConfigType(in)
}

func (this *Core) SetConfigName(in string) {
	this.getCore().SetConfigName(in)
}

func (this *Core) AddConfigPath(in string) {
	this.getCore().AddConfigPath(in)
}

func (this *Core) getCore() *viper.Viper {
	return this.core
}

func (this *Core) SetDefault(key string, value interface{}) CoreInterface {
	this.getCore().SetDefault(key, value)
	return this
}

func (this *Core) WriteConfig() error {
	return this.core.SafeWriteConfig()
}

func (this *Core) WriteConfigAs(path string) error {
	return this.core.SafeWriteConfigAs(path)
}

func (this *Core) ReadConfig(data []byte) (*viper.Viper, error) {
	if data != nil {
		if err := this.getCore().ReadConfig(bytes.NewBuffer(data)); err == nil {
			return this.getCore(), nil
		} else {
			return nil, err
		}
	} else {
		if err := this.getCore().ReadInConfig(); err == nil {
			return this.getCore(), nil
		} else {
			return nil, err
		}
	}
}
