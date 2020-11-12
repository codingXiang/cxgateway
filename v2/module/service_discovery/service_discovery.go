package service_discovery

import (
	"github.com/codingXiang/configer/v2"
	"github.com/codingXiang/service-discovery/discovery"
	"github.com/codingXiang/service-discovery/info"
	"github.com/codingXiang/service-discovery/register"
	"strings"
)

var ServiceDiscovery *discovery.ServiceDiscovery
var Register *register.ServiceRegister

const (
	_etcd      = "etcd"
	_endpoints = "endpoints"
)

const (
	_register = "register"
	_prefix   = "prefix"
	_key      = "key"
	_name     = "name"
	_addr     = "addr"
	_leave    = "leave"
)

func Init(_type configer.FileType, configName string, path ...string) {
	if config, err := configer.NewCore(_type, configName, path...).ReadConfig(); err == nil {
		endpoints := strings.Split(config.GetString(configer.GetConfigPath(_etcd, _endpoints)), ",")
		lease := config.GetInt64(configer.GetConfigPath(_register, _leave))
		prefix := config.GetString(configer.GetConfigPath(_register, _prefix))
		name := config.GetString(configer.GetConfigPath(_register, _name))
		key := config.GetString(configer.GetConfigPath(_register, _key))
		addr := config.GetString(configer.GetConfigPath(_register, _addr))

		ServiceDiscovery = discovery.New(endpoints)

		if r, err := register.New(endpoints, info.New(prefix, key, name, addr), lease); err == nil {
			Register = r
		} else {

		}
	}
}

func StartWatch(prefix string) {
	ServiceDiscovery.WatchService(prefix)
}
