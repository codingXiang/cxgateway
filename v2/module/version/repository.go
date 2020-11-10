package version

import "github.com/codingXiang/cxgateway/v2/model"

type Repository interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
