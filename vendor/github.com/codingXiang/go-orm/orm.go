package orm

import (
	"fmt"
	"github.com/codingXiang/configer"
	"github.com/codingXiang/go-logger"
	. "github.com/codingXiang/go-orm/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    // mysql
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgresql
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

//Orm
type (
	OrmInterface interface {
		CloseDB()
		GetTableName(value interface{}) string
		CheckTable(migrate bool, value interface{}) error
		GetInstance() *gorm.DB
		SetInstance(db *gorm.DB)
		CheckVersion() error
		Upgrade(tables ...interface{}) error
	}
	Orm struct {
		db         *gorm.DB
		configName string
		version    *Version
	}
)

var (
	DatabaseORM OrmInterface
)

//NewOrm : 新增 ORM 實例
func NewOrm(configName string, core configer.CoreInterface) (OrmInterface, error) {
	var o = &Orm{
		configName: configName,
	}
	return o.init(core)
}

//init : 初始化 ORM
func (this *Orm) init(core configer.CoreInterface) (OrmInterface, error) {
	var (
		data *viper.Viper
		err  error
	)
	if configer.Config == nil {
		//初始化 configer
		configer.Config = configer.NewConfiger()
	}
	//加入 config
	configer.Config.AddCore(this.configName, core)
	//讀取 config
	if data, err = configer.Config.GetCore(this.configName).ReadConfig(nil); err == nil {
		var (
			logMode     = data.GetBool("database.logMode")
			tablePrefix = data.GetString("database.tablePrefix")
			version     = data.GetString("database.version")
		)
		//設定資料庫型態 (MySQL, PostgreSQL) 與連線資訊
		logger.Log.Debug("setup database type")
		err = this.setDatabaseType(data)

		//設定資料庫參數
		logger.Log.Debug("setup database config")
		this.setDbConfig(data)

		//設定是否開啟 Log 模式
		logger.Log.Debug("setup log mode =", logMode)
		if logMode {
			this.GetInstance().LogMode(true)
			this.SetInstance(this.GetInstance().Debug())
		} else {
			this.GetInstance().LogMode(false)
		}

		//設定預設 Table 前綴字
		logger.Log.Debug("setup table name prefix", tablePrefix)
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return tablePrefix + defaultTableName
		}

		//設定版本控制 Schema
		this.version = &Version{Version: version}
		logger.Log.Debug("setup version", this.version.GetVersion())
		err = this.setVersion(this.version)
	}

	return this, err
}

func (this *Orm) CloseDB() {
	logger.Log.Debug("close database connection")
	defer this.GetInstance().Close()
}

//GetTableName : 透過傳入 struct 回傳 table 名稱
func (this *Orm) GetTableName(tb interface{}) string {
	return this.GetInstance().NewScope(tb).TableName()
}

//CheckTable : 檢查 Table 是否存在，不存在建立並回傳 false, 反之回傳 true
func (this *Orm) CheckTable(migrate bool, value interface{}) error {
	var (
		err error
		tx  = this.GetInstance().Begin()
	)
	if !this.GetInstance().HasTable(value) {
		if err = tx.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").CreateTable(value).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else {
		if migrate {
			if err = tx.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(value).Error; err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()
		}
	}
	return nil
}

func (orm *Orm) GetInstance() *gorm.DB {
	return orm.db
}

func (this *Orm) SetInstance(db *gorm.DB) {
	this.db = db
}

//setDatabaseType : 設定資料庫型態
func (this *Orm) setDatabaseType(config *viper.Viper) error {
	var (
		err      error
		url      = config.GetString("database.url")
		port     = config.GetInt("database.port")
		dbName   = config.GetString("database.name")
		username = config.GetString("database.username")
		password = config.GetString("database.password")
		_type    = config.GetString("database.type")
	)
	logger.Log.Debug("set database type = ", _type)
	switch _type {
	case "mysql":
		var connectStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
			username, password, url, port, dbName)
		logger.Log.Debug("connection string = ", connectStr)
		this.db, err = gorm.Open(_type, connectStr)
		if err != nil {
			logger.Log.Error("connect to database error", err)
			return err
		}
		break
	case "postgres":
		var connectStr = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			url, port, dbName, username, password)
		logger.Log.Debug("connection string = ", connectStr)
		this.db, err = gorm.Open(_type, connectStr)
		if err != nil {
			logger.Log.Error("connect to database error", err)
			return err
		}
		break
	}
	return nil
}

// 設定資料庫組態
func (this *Orm) setDbConfig(config *viper.Viper) {
	var (
		maxOpenConns = config.GetInt("database.maxOpenConns")
		maxIdleConns = config.GetInt("database.maxIdleConns")
		maxLifeTime  = config.GetInt("database.maxLifeTime")
	)
	logger.Log.Debug("set max idle connections", maxIdleConns)
	this.GetInstance().DB().SetMaxIdleConns(maxIdleConns)
	logger.Log.Debug("set max open connections", maxOpenConns)
	this.GetInstance().DB().SetMaxOpenConns(maxOpenConns)
	logger.Log.Debug("set connection max life time", maxLifeTime)
	this.GetInstance().DB().SetConnMaxLifetime(time.Duration(maxLifeTime))
}

//設定資料庫版本
func (this *Orm) setVersion(version *Version) error {
	var (
		err error
		vs  []*Version
		tx  = this.GetInstance().Begin()
	)

	if err = tx.Model(&Version{}).Find(&vs).Error; err != nil {
		logger.Log.Debug("not found version table")
		if err = this.CheckTable(false, &version); err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Model(&Version{}).Create(&version).Error; err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
		}
		return nil
	} else if len(vs) < 1 {
		logger.Log.Debug("not found version record")
		if err := tx.Model(&Version{}).Create(&version).Error; err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
		}
		return nil
	}
	return err
}

func (this *Orm) CheckVersion() error {
	var (
		err     error
		vs      []*Version
		version = this.version.GetVersion()
		tx      = this.GetInstance().Begin()
	)
	/*
		檢查是否可以更新
	*/
	if err = tx.Model(&Version{}).Find(&vs).Error; err != nil {
		logger.Log.Debug("not found version table")
		if err = this.CheckTable(false, &version); err != nil {
			return err
		}
		if err := tx.Model(&Version{}).Create(&version).Error; err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
		}
	} else if len(vs) < 1 {
		logger.Log.Debug("not found version record")
		if err := tx.Model(&Version{}).Create(&version).Error; err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
		}
	}
	var oldVersion = vs[0]
	logger.Log.Debug("found old version", oldVersion.GetVersion())
	if oldVersion.GetVersion() != this.version.GetVersion() {
		var (
			ov int
			nv int
		)
		//判斷版本是否高於現有版本
		if ov, err = this.transformVersion(oldVersion.GetVersion()); err != nil {
			logger.Log.Error("transform old version error", err)
			return err
		}
		if nv, err = this.transformVersion(this.version.GetVersion()); err != nil {
			logger.Log.Error("transform new version error", err)
			return err
		}
		if nv > ov {
			return nil
		} else {
			return errors.New("v" + this.version.GetVersion() + " is not higher than v" + oldVersion.GetVersion())
		}
	} else {
		return errors.New("v" + this.version.GetVersion() + " is the newest version")
	}
}

func (this *Orm) Upgrade(tables ...interface{}) error {
	if err := this.CheckVersion(); err == nil {
		for _, table := range tables {
			if err := this.CheckTable(true, table); err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}

////loadSQLFile 讀取更新的 SQL 檔案
//func (this *Orm) loadSQLFile(version string) ([]byte, error) {
//	var (
//		file = version + ".sql"
//		sql  []byte
//		err  error
//	)
//	logger.Log.Debug("read sql file", file)
//	if sql, err = ioutil.ReadFile(this.config.GetUpgradeFilePath() + file); err != nil {
//		return nil, err
//	}
//	return sql, nil
//}

//translformVersion 轉換版本權重
func (this *Orm) transformVersion(version string) (int, error) {
	var (
		tmp        []string
		err        error
		v1, v2, v3 int
		result     int = 0
	)
	tmp = strings.Split(version, ".")
	if v1, err = strconv.Atoi(tmp[0]); err != nil {
		return 0, err
	}
	if v1, err = strconv.Atoi(tmp[1]); err != nil {
		return 0, err
	}
	if v1, err = strconv.Atoi(tmp[2]); err != nil {
		return 0, err
	}

	result = v1*100 + v2*10 + v3

	return result, err
}
