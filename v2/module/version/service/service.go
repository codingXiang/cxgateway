package service

import (
	"github.com/codingXiang/cxgateway/v2/model"
	"github.com/codingXiang/cxgateway/v2/module/version"
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
