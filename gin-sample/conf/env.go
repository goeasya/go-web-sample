package conf

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EnvDBType        = "DB_TYPE"
	EnvDBConnInfo    = "DB_CONN_INFO"
	EnvDBAutoMigrate = "DB_AUTO_MIGRATE"
	EnvSecretKey     = "SECRET_KEY"
	EnvApiAddr       = "API_ADDR"
	EnvLogLevel      = "LOG_LEVEL"
	EnvLogEncoder    = "LOG_ENCODER"
)

func init() {
	RegisterServer("env", makeEnvServer)
}

// EnvServer
type EnvServer struct{}

func makeEnvServer() (Interface, error) {
	return &EnvServer{}, nil
}

func (s *EnvServer) LoadConfig() (*Config, error) {
	// get config
	apiCfg, err := s.getApiConfig()
	if err != nil {
		return nil, err
	}
	dbCfg, err := s.getDataBaseConfig()
	if err != nil {
		return nil, err
	}
	globalCfg, err := s.getGlobalConfig()
	if err != nil {
		return nil, err
	}

	storeCfg := &Config{
		GlobalCfg: globalCfg,
		ApiCfg:    apiCfg,
		DBCfg:     dbCfg,
	}
	return storeCfg, nil
}

func (s *EnvServer) getApiConfig() (*ApiConfig, error) {
	apiAddr := os.Getenv(EnvApiAddr)
	if apiAddr == "" {
		return nil, ErrInvalidApiAddr
	}
	apiCfg := &ApiConfig{
		APIAddr: apiAddr,
	}
	return apiCfg, nil
}

func (s *EnvServer) getDataBaseConfig() (*DatabaseConfig, error) {
	dbType := os.Getenv(EnvDBType)
	dbConnInfo := os.Getenv(EnvDBConnInfo)
	dbAutoMigrate, _ := strconv.ParseBool(os.Getenv(EnvDBAutoMigrate))
	if dbType == "" || dbConnInfo == "" {
		return nil, fmt.Errorf("invalid db config")
	}
	dbCfg := &DatabaseConfig{
		Type:        dbType,
		ConnInfo:    dbConnInfo,
		AutoMigrate: dbAutoMigrate,
	}
	return dbCfg, nil
}

func (s *EnvServer) getGlobalConfig() (*GlobalConfig, error) {
	secretKey := os.Getenv(EnvSecretKey)
	if secretKey == "" {
		return nil, ErrInvalidSecretKey
	}

	logLevel := GetLogLevel(os.Getenv(EnvLogLevel))
	logEncoder := GetLogEncoder(os.Getenv(EnvLogEncoder))
	globalCfg := &GlobalConfig{
		SecretKey:  secretKey,
		LogLevel:   logLevel,
		LogEncoder: logEncoder,
	}

	return globalCfg, nil
}
