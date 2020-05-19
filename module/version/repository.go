package version

import "github.com/codingXiang/go-orm/model"

type Repository interface {
	GetServerVersion() (*model.Version, error)
	CheckVersion() error
	Upgrade() error
}
