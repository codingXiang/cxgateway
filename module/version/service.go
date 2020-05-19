package version

import "github.com/codingXiang/go-orm/model"

type Service interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
