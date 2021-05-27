package conf

type Interface interface {
	// load all config. include db, log, arena...
	LoadConfig() (*Config, error)
}
