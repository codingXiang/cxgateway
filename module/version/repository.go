package version

import "github.com/codingXiang/cxgateway/model"

type Repository interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
