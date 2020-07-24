package version

import "github.com/codingXiang/cxgateway/model"

type Service interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
