package factory

import (
	"gin-sample/conf"
	dbinterface "gin-sample/db/interface"
)

var dbProviderMap = map[string]Factory{}

type Factory func(cfg *conf.DatabaseConfig) (dbinterface.Database, error)

func RegisterDbFactory(name string, f Factory) {
	dbProviderMap[name] = f
}

func GetDatabaseFactory(dbType string) (Factory, bool) {
	f, ok := dbProviderMap[dbType]
	return f, ok
}
