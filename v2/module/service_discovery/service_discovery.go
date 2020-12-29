package service_discovery

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/service-discovery/discovery"
	"github.com/codingXiang/service-discovery/info"
	"github.com/codingXiang/service-discovery/register"
	"github.com/codingXiang/service-discovery/util"
	"strings"
)

var ServiceDiscovery *discovery.ServiceDiscovery
var Register *register.ServiceRegister

const (
	_etcd      = "etcd"
	_endpoints = "endpoints"
	_username  = "username"
	_password  = "password"
)

const (
	_register = "register"
	_prefix   = "prefix"
	_key      = "key"
	_name     = "name"
	_addr     = "addr"
	_leave    = "leave"
)

func Init(prefix string, _type configer.FileType, configName string, path ...string) {
	c := configer.NewCore(_type, configName, path...)
	c.SetAutomaticEnv(prefix, ".", "_")
	if config, err := c.ReadConfig(); err == nil {
		endpoints := strings.Split(config.GetString(configer.GetConfigPath(_etcd, _endpoints)), ",")
		username := config.GetString(configer.GetConfigPath(_etcd, _username))
		password := config.GetString(configer.GetConfigPath(_etcd, _password))
		lease := config.GetInt64(configer.GetConfigPath(_register, _leave))
		prefix := config.GetString(configer.GetConfigPath(_register, _prefix))
		name := config.GetString(configer.GetConfigPath(_register, _name))
		key := config.GetString(configer.GetConfigPath(_register, _key))
		addr := config.GetString(configer.GetConfigPath(_register, _addr))

		auth := &util.ETCDAuth{
			Endpoints: endpoints,
			Username:  username,
			Password:  password,
		}

		ServiceDiscovery = discovery.New(auth)
		healthCheck(auth, info.New(prefix, key, name, addr), lease)
	}
}

func healthCheck(auth *util.ETCDAuth, i *info.ServiceInfo, lease int64) {
	if r, err := register.New(auth, i, lease); err == nil {
		Register = r
	} else {
		logger.Log.Error(err)
	}
}

func StartWatch(prefix string) error{
	defer Register.Close()
	defer ServiceDiscovery.Close()
	return ServiceDiscovery.WatchService(prefix)
}

func StartWatchBackground(prefix string) {
	go StartWatch(prefix)
}
