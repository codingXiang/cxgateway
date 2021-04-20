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
	github.com/codingXiang/gogo-i18n v1.0.2-0.20200417093325-c191114c00c4
	github.com/dapr/go-sdk v1.1.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/ugorji/go v1.2.5 // indirect
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc // indirect
	golang.org/x/sys v0.0.0-20210419170143-37df388d1f33 // indirect
	golang.org/x/text v0.3.6
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
