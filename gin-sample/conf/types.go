package conf

import "errors"

const (
	LogLevelDebug   = "debug"
	LogLevelInfo    = "info"
	LogLevelWarn    = "warn"
	LogLevelError   = "error"
	LogLevelFatal   = "fatal"
	LogLevelDefault = "info"

	LogConsoleEncoder = "console"
	LogJsonEncoder    = "json"
)

var (
	ErrInvalidSecretKey = errors.New("invalid secret key")
	ErrInvalidApiAddr   = errors.New("invalid api addr")
)

type (
	Config struct {
		GlobalCfg *GlobalConfig
		ApiCfg    *ApiConfig
		DBCfg     *DatabaseConfig
	}

	GlobalConfig struct {
		LogLevel   string `json:"loglevel,default=info,options=debug|info|warn|error|fatal"`
		LogEncoder string `json:"logEncoder,default=console,options=console|json"`
		SecretKey  string
	}

	ApiConfig struct {
		APIAddr string
	}

	DatabaseConfig struct {
		Type        string
		ConnInfo    string
		AutoMigrate bool
	}
)

func GetLogLevel(level string) string {
	if level == "" {
		return LogLevelDefault
	}

	logLevelMap := map[string]string{
		LogLevelDebug: LogLevelDebug,
		LogLevelInfo:  LogLevelInfo,
		LogLevelWarn:  LogLevelWarn,
		LogLevelError: LogLevelError,
		LogLevelFatal: LogLevelFatal,
	}
	if l, ok := logLevelMap[level]; ok {
		return l
	}
	return LogLevelDefault
}

func GetLogEncoder(encoder string) string {
	if encoder == LogJsonEncoder {
		return LogJsonEncoder
	}
	return LogConsoleEncoder
}
