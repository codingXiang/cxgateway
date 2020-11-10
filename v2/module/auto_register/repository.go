package auto_register

import (
	"github.com/codingXiang/cxgateway/v2/model"
)

//Repository 用於與資料庫進行存取的封裝方法
//go:generate mockgen -destination mock/mock_repository.go -package mock -source repository.go
type Repository interface {
	/*
	   以下宣告 Repository 方法
	*/
	GetConfig(key string) (string, error)
	Register(data *model.ServiceRegister) (*model.ServiceRegister, error)
	Initial() error
}
