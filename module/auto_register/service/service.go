package service

import (
	"github.com/codingXiang/cxgateway/model"
	"github.com/codingXiang/cxgateway/module/auto_register"
)

type AutoRegisteredService struct {
	Repo auto_register.Repository
}

func NewAutoRegisteredService(repo auto_register.Repository) auto_register.Service {
	return &AutoRegisteredService{Repo: repo}
}

func (a *AutoRegisteredService) GetConfig(key string) (*model.ServiceRegister, error) {
	return a.Repo.GetConfig(key)
}

func (a *AutoRegisteredService) Register(data *model.ServiceRegister) (*model.ServiceRegister, error) {
	return a.Repo.Register(data)
}

func (a *AutoRegisteredService) Initial() error {
	return a.Repo.Initial()
}
