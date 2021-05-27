package mysql

import (
	"errors"
	"sync"
	"time"

	"gin-sample/conf"
	"gin-sample/db/factory"
	dbinterface "gin-sample/db/interface"
	"gin-sample/db/mysql/impl"
	"gin-sample/global"
	"gin-sample/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	factory.RegisterDbFactory("mysql", CreateMysqlManager)
}

type Manager struct {
	autoMigrate bool
	db          *gorm.DB
	initOnce    sync.Once
	models      []model.Interface
}

func CreateMysqlManager(cfg *conf.DatabaseConfig) (dbinterface.Database, error) {
	if cfg.Type != "mysql" {
		return nil, errors.New("db driver not support")
	}

	dsn := cfg.ConnInfo + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Manager{
		db:          db,
		initOnce:    sync.Once{},
		autoMigrate: cfg.AutoMigrate,
	}, nil
}

func (m *Manager) InitDataBase() {
	m.initModels()
	if m.autoMigrate {
		m.checkAndInitTable()
	}
}

func (m *Manager) initModels() {
	m.models = append(m.models, &model.UserInfo{})
}

func (m *Manager) checkAndInitTable() {
	m.initOnce.Do(func() {
		for _, mo := range m.models {
			if !m.db.Migrator().HasTable(mo) {
				err := m.db.Migrator().CreateTable(mo)
				if err != nil {
					global.Logger.Errorf("auto create table %s error", mo.TableName())
				} else {
					global.Logger.Debugf("auto create table %s success", mo.TableName())
				}
			}
		}
	})
}

func (m *Manager) Users() dbinterface.UserInterface {
	return &impl.UserInfoImpl{
		DB: m.db,
	}
}
