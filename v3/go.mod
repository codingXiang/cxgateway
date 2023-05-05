module github.com/codingXiang/cxgateway/v3

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/astaxie/beego v1.12.3
	github.com/codingXiang/configer/v2 v2.0.3
	github.com/codingXiang/cxgateway/v2 v2.0.0-20201229021159-426b9cb8f2d0
	github.com/codingXiang/go-logger v1.0.2
	github.com/codingXiang/go-logger/v2 v2.0.5
	github.com/codingXiang/go-orm/v2 v2.0.9
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.9.0
	github.com/satori/go.uuid v1.2.1-0.20180404165556-75cca531ea76
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/valyala/fasthttp v1.24.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.21.10
)
