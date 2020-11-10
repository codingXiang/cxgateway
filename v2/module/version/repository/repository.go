package repository

import (
	"github.com/codingXiang/cxgateway/v2/model"
	"github.com/codingXiang/cxgateway/v2/module/version"
	"github.com/codingXiang/go-orm"
	"io/ioutil"
	"strings"
)

var VERSION_CONTROL = "./version_control"

type VersionRepository struct {
	db     orm.OrmInterface
	redis  orm.RedisClientInterface
	tables []interface{}
}

func NewVersionRepository(db orm.OrmInterface, redis orm.RedisClientInterface, tables ...interface{}) version.Repository {
	return &VersionRepository{
		db:     db,
		redis:  redis,
		tables: tables,
	}
}
func (this *VersionRepository) GetServerVersion() (*model.Version, error) {
	var (
		version      = new(model.Version)
		appVersion   string
		buildVersion string
	)
	appVersionTmp, err := ioutil.ReadFile(VERSION_CONTROL + "/APP_VERSION")
	if err != nil {
		return nil, err
	} else {
		appVersion = strings.ReplaceAll(string(appVersionTmp), "\n", "")
	}
	buildVersionTmp, err := ioutil.ReadFile(VERSION_CONTROL + "/BUILD")
	if err != nil {
		return nil, err
	} else {
		buildVersion = strings.ReplaceAll(string(buildVersionTmp), "\n", "")
	}
	version.ServerVersion = "v" + string(appVersion) + "." + string(buildVersion)
	// 抓取資料庫版本
	if this.db != nil {
		version.DatabaseVersion = this.db.ShowVersion()
	}
	// 抓取 redis 版本
	if this.redis != nil {
		if info := this.redis.Info("server")["server"]; info != nil {
			if v := info["redis_version"]; v != nil || v != "" {
				version.RedisVersion = v.(string)
			}
		}
	}
	return version, err
}

func (this *VersionRepository) CheckVersion() error {
	return this.db.CheckVersion()
}
func (this *VersionRepository) Upgrade() error {
	return this.db.Upgrade(this.tables)
}
