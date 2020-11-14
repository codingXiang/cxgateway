module github.com/codingXiang/cxgateway/v2

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/astaxie/beego v1.12.1
	github.com/codingXiang/configer v1.0.2-0.20200513072245-ec8070de9a16
	github.com/codingXiang/configer/v2 v2.0.3
	github.com/codingXiang/go-logger v1.0.2
	github.com/codingXiang/go-logger/v2 v2.0.5
	github.com/codingXiang/go-orm v1.0.7
	github.com/codingXiang/go-orm/v2 v2.0.0
	github.com/codingXiang/service-discovery v1.0.0
	github.com/codingXiang/gogo-i18n v1.0.2-0.20200417093325-c191114c00c4
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.2
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.15
	github.com/lib/pq v1.7.1 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.3.2 // indirect
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sys v0.0.0-20200722175500-76b94024e4b6 // indirect
	golang.org/x/text v0.3.3
	google.golang.org/grpc v1.27.0
	gopkg.in/ini.v1 v1.57.0 // indirect
)
