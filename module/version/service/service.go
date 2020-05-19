package service

import (
	"github.com/codingXiang/cxgateway/module/version"
	"github.com/codingXiang/go-orm/model"
)

type VersionService struct {
	repo version.Repository
}

func NewVersionService(repo version.Repository) version.Service {
	return &VersionService{
		repo: repo,
	}
}
func (this *VersionService) GetServerVersion() (*model.Version, error) {
	return this.repo.GetServerVersion()
}

func (this *VersionService) CheckVersion() error {
	return this.repo.CheckVersion()
}

func (this *VersionService) Upgrade() error {
	return this.repo.Upgrade()
}
