package server

import (
	"fmt"
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/cxgateway/v2/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	Gateway *Server
)

type Server struct {
	config     *viper.Viper
	engine     *gin.Engine
	api        *gin.RouterGroup
	server     *http.Server
	appId      string
	uploadPath string
}

func New(engine *gin.Engine, config *viper.Viper) *Server {
	s := new(Server)
	//設定 gin 啟動模式
	gin.SetMode(config.GetString(configer.GetConfigPath(Application, Mode)))
	//設定 server config
	s.config = config
	//設定 server engine
	if engine == nil {
		s.engine = gin.Default()
	} else {
		s.engine = engine
	}

	var (
		appId        = config.GetString(configer.GetConfigPath(Application, AppId))
		port         = config.GetInt(configer.GetConfigPath(Application, Port))           //伺服器的 port
		writeTimeout = config.GetInt(configer.GetConfigPath(Application, Timeout, Write)) //伺服器的寫入超時時間
		readTimeout  = config.GetInt(configer.GetConfigPath(Application, Timeout, Read))  //伺服器讀取超時時間
	)
	// 設定 http server
	s.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        s.engine,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.appId = appId
	s.uploadPath = config.GetString(configer.GetConfigPath(Application, UploadPath))
	return s
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) GetApiRoute() *gin.RouterGroup {
	return s.api
}

func (s *Server) GetServer() *http.Server {
	return s.server
}

func (s *Server) GetAppID() string {
	return s.appId
}

//Run 運行 Server
func (s *Server) Run() {
	s.GetServer().ListenAndServe()
}

//Stop 停止 Server
func (s *Server) Stop() {
	s.GetServer().Close()
}

func (s *Server) Use(handle ...middleware.Object) *gin.Engine {
	for _, h := range handle {
		if h.GetConfig() == nil {
			h.SetConfig(s.GetConfig())
		}
		s.GetEngine().Use(h.Handle())
	}
	// 設定 api routing
	s.api = s.GetEngine().Group(s.config.GetString(configer.GetConfigPath(Application, BaseRoute)))

	return s.GetEngine()
}

func (s *Server) AddModule(modules ...HttpModule) error {
	for _, m := range modules {
		m.SetGateway(s)
		m.Setup()
	}
	return nil
}

func (s *Server) GetConfig() *viper.Viper {
	return s.config
}
