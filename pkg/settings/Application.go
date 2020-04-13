package settings

import (
	"github.com/codingXiang/go-logger"
)

type (
	//ApplicationInterface 應用參數介面
	ApplicationInterface interface {
		GetPort() int
		GetMode() string
		GetLog() logger.Logger
		GetTimeout() TimeoutInterface
	}
	//TimeoutInterface 超時介面參數
	TimeoutInterface interface {
		GetRead() int
		GetWrite() int
	}
	//Application : 應用程式相關參數
	Application struct {
		Port     int           `yaml:"port"`     //運行開放port
		Mode     string        `yaml:"mode"`     //運行模式（有release、test 與 debug）
		Log      logger.Logger `yaml:"log"`      //Log 配置
		AppToken string        `yaml:"appToken"` //應用 id
		Timeout  *Timeout      `yaml:"timeout"`  //超時參數
	}
	//Timeout : 超時參數
	Timeout struct {
		Read  int `yaml:"read"`
		Write int `yaml:"write"`
	}
)

//GetPort 取得 port 號
func (a *Application) GetPort() int {
	return a.Port
}

//GetMode 取得運行的模式
func (a *Application) GetMode() string {
	return a.Mode
}

//GetLog 取得 log 設定
func (a *Application) GetLog() logger.Logger {
	return a.Log
}

//GetTimeout 取得 Timeout
func (a *Application) GetTimeout() TimeoutInterface {
	return a.Timeout
}

//GetRead 取得讀取 timeout 時間
func (t *Timeout) GetRead() int {
	return t.Read
}

//GetWrite 取得寫入 timeout 時間
func (t *Timeout) GetWrite() int {
	return t.Write
}
