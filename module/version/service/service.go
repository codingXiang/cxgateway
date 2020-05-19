package service

import (
	"github.com/codingXiang/cxgateway/module/version"
	"github.com/codingXiang/cxgateway/module/version/repository"
	"github.com/codingXiang/go-orm/model"
)

type VersionService struct {
	repo repository.VersionRepository
}

func NewVersionService(repo repository.VersionRepository) version.Service {
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
