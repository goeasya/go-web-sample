package conf

import (
	"fmt"
	"os"
)

var configMap = map[string]factory{}

type factory func() (Interface, error)

func RegisterServer(name string, f factory) {
	configMap[name] = f
}

func MustLoad() *Config {
	server, err := constructServer()
	if err != nil {
		fmt.Printf("config construct server error: %v", err.Error())
		os.Exit(1)
	}

	c, err := server.LoadConfig()
	if err != nil {
		fmt.Printf("load config error: %v", err.Error())
		os.Exit(1)
	}
	return c
}

func constructServer() (Interface, error) {
	tp := loadConfigType()
	factory, ok := configMap[tp]
	if !ok {
		return nil, fmt.Errorf("unknown type %s", tp)
	}

	return factory()
}

// loadConfigType, get CONFIG_TYPE (env, zookeeper, file)
func loadConfigType() string {
	tp := os.Getenv("CONFIG_TYPE")
	if tp == "" {
		tp = "env"
	}
	fmt.Println("config type is " + tp)
	return tp
}
