package global

import "gin-sample/conf"

var secretKey string

func Init(cfg *conf.GlobalConfig) error {
	secretKey = cfg.SecretKey
	return initLogger(cfg.LogLevel, cfg.LogEncoder)
}

func GetSecretKey() []byte {
	return []byte(secretKey)
}
