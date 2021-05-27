package db

import (
	"fmt"

	"gin-sample/conf"
	"gin-sample/db/factory"
	dbinterface "gin-sample/db/interface"
	_ "gin-sample/db/mysql"
)

var defaultDBManager dbinterface.Database

func Init(cfg *conf.DatabaseConfig) error {
	var err error
	defaultDBManager, err = NewDBManager(cfg)
	return err
}

func NewDBManager(cfg *conf.DatabaseConfig) (dbinterface.Database, error) {
	newManagerFunc, ok := factory.GetDatabaseFactory(cfg.Type)
	if !ok {
		return nil, fmt.Errorf("unsport database type %s", cfg.Type)
	}

	manager, err := newManagerFunc(cfg)
	if err != nil {
		return nil, fmt.Errorf("new database manager type %s error %v", cfg.Type, err.Error())
	}
	manager.InitDataBase()
	return manager, nil
}

func Manager() dbinterface.Database {
	return defaultDBManager
}
