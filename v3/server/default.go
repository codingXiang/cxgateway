package server

import (
	"github.com/codingXiang/configer/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const TimeFormat = "2006/05/26:15:04:05 +0800"

const (
	Application = "application"
	Port        = "port"
	Timeout     = "timeout"
	Read        = "read"
	Write       = "write"
	Mode        = "mode"
	Enable      = "enable"
	Version     = "version"
	Key         = "key"
	Value       = "value"
	AppId       = "appId"
	BaseRoute   = "apiBaseRoute"
	UploadPath  = "uploadPath"
)

const ()

const (
	I18n            = "i18n"
	DefaultLanguage = "defaultLanguage"
	File            = "file"
	Path            = "Path"
	Type            = "type"
)

const (
	Default_       = "default"
	Cors           = "cors"
	AllowAllOrigin = "allowAllOrigin"
	AllowOrigins   = "allowOrigins"
	AllowHeaders   = "allowHeaders"
	AllowMethods   = "allowMethods"
)

const (
	LOG      = "log"
	Format   = "format"
	Level    = "level"
	MaxAge   = "maxAge"
	Filename = "filename"
)

func DefaultConfig() *viper.Viper {
	config := viper.New()
	config.Set(Cors, map[string]interface{}{
		Default_:       true,
		AllowAllOrigin: true,
		AllowOrigins:   "*",
		AllowHeaders:   "*",
		AllowMethods:   "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	})
	config.Set(Application, map[string]interface{}{
		Timeout: map[string]interface{}{
			Read:  1000,
			Write: 1000,
		},
		Port:       9999,
		UploadPath: "./upload",
		Mode:       gin.DebugMode,
		AppId:      "test",
		BaseRoute:  "/api",
		Version: map[string]interface{}{
			Enable: false,
		},
	})
	config.Set(configer.GetConfigPath(I18n, File, Path), "./i18n")
	config.Set(configer.GetConfigPath(I18n, File, Type), configer.YAML.String())
	config.Set(configer.GetConfigPath(I18n, Enable), true)
	config.Set(configer.GetConfigPath(I18n, DefaultLanguage), "zh_Hant")

	config.Set(LOG, map[string]interface{}{
		Level:    gin.DebugMode,
		Format:   configer.JSON.String(),
		MaxAge:   7,
		Path:     "log",
		Filename: "test.log",
	})
	return config
}

func Default() *Server {
	s := New(nil, DefaultConfig())
	return s
}