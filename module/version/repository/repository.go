package repository

import (
	"github.com/codingXiang/cxgateway/module/version"
	"github.com/codingXiang/go-orm"
	"github.com/codingXiang/go-orm/model"
)

type VersionRepository struct {
	db     orm.OrmInterface
	tables []interface{}
}

func NewVersionRepository(db orm.OrmInterface) version.Repository {
	return &VersionRepository{
		db: db,
	}
}
func (this *VersionRepository) GetServerVersion() (*model.Version, error) {
	version := new(model.Version)
	err := this.db.GetInstance().First(&version).Error
	return version, err
}

func (this *VersionRepository) CheckVersion() error {
	return this.db.CheckVersion()
}
func (this *VersionRepository) Upgrade() error {
	return this.db.Upgrade(this.tables)
}
