package server

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

var DefaultConfig = []byte(`cors:
  allowAllOrigin: true
  allowOrigins: "*"
  allowHeaders: "*"
  allowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS"
application:
  timeout:
    read: 1000
    write: 1000
  port: 9999
  uploadPath: "./upload"
  mode: "debug"
  appId: ""
  appToken: ""
  apiBaseRoute: "/api"
  version:
    enable: true
    key: app-version
    value: 1.0.0
i18n:
  enable: true
  defaultLanguage: "zh_Hant"
  file:
    path: "./i18n"
    type: "yaml"
log:
  level: "debug"
  format: "json"
  path: "log"
  filename: "gateway.log"`)
