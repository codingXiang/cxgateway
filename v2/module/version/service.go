package version

import "github.com/codingXiang/cxgateway/v2/model"

type Service interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
