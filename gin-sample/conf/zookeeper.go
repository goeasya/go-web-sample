package conf

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
)

const (
	EnvConfigZKServer = "CONFIG_ZK_SERVER"
	EnvConfigZKNode   = "CONFIG_ZK_NODE"

	PropertyDBType        = "db_type"
	PropertyDBConnInfo    = "db_conn_info"
	PropertyDBAutoMigrate = "db_auto_migrate"
	PropertySecretKey     = "secret_key"
	PropertyApiAddr       = "api_addr"
	// options: fatal|error|warn|info|debug. default: info
	PropertyLoglevel = "loglevel"
	// options: json|console. default: console
	PropertyLogEncoder = "log_encoder"
)

func init() {
	RegisterServer("zookeeper", makeZooKeeperServer)
}

// ZooKeeperServer
type ZooKeeperServer struct {
	Server string
	Node   string
}

func makeZooKeeperServer() (Interface, error) {
	server := os.Getenv(EnvConfigZKServer)
	node := os.Getenv(EnvConfigZKNode)
	if server == "" || node == "" {
		return nil, fmt.Errorf("makeZooKeeperServer error: CONFIG_ZK_SERVER is nil or CONFIG_ZK_NODE is nil")
	}
	s := &ZooKeeperServer{
		Server: server,
		Node:   node,
	}
	return s, nil
}

func (s *ZooKeeperServer) LoadConfig() (*Config, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	basePath := "/" + s.Node + "/"
	// get config
	dbCfg, err := s.getDatabaseConfig(basePath, conn)
	if err != nil {
		return nil, err
	}
	apiCfg, err := s.getApiConfig(basePath, conn)
	if err != nil {
		return nil, err
	}
	globalCfg, err := s.getGlobalConfig(basePath, conn)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		GlobalCfg: globalCfg,
		ApiCfg:    apiCfg,
		DBCfg:     dbCfg,
	}
	return cfg, nil
}

func (s *ZooKeeperServer) connect() (*zk.Conn, error) {
	serverList := strings.Split(s.Server, ",")
	if len(serverList) == 0 {
		return nil, fmt.Errorf("invalid server format: %s", s.Server)
	}

	conn, _, err := zk.Connect(serverList, 10*time.Second)
	return conn, err
}

func (s *ZooKeeperServer) getNodeProperty() (map[string]string, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	children, _, err := conn.Children("/" + s.Node)
	if err != nil {
		return nil, err
	}

	propertyMap := make(map[string]string, 0)
	for _, property := range children {
		value, _, err := conn.Get("/" + s.Node + "/" + property)
		if err != nil {
			return nil, err
		}
		propertyMap[property] = string(value)
	}

	return propertyMap, nil
}

func (s *ZooKeeperServer) getDatabaseConfig(basePath string, conn *zk.Conn) (*DatabaseConfig, error) {
	tp, _, err := conn.Get(basePath + PropertyDBType)
	if err != nil {
		return nil, fmt.Errorf("get %s error: %v", PropertyDBType, err)
	}

	connInfo, _, err := conn.Get(basePath + PropertyDBConnInfo)
	if err != nil {
		return nil, fmt.Errorf("get %s error: %v", PropertyDBConnInfo, err)
	}

	autoMigrate, _, err := conn.Get(basePath + PropertyDBAutoMigrate)
	if err != nil {
		return nil, fmt.Errorf("get %s error: %v", PropertyDBAutoMigrate, err)
	}
	migrate, err := strconv.ParseBool(string(autoMigrate))
	if err != nil {
		return nil, fmt.Errorf("conver %s bool error: %v", PropertyDBAutoMigrate, err)
	}

	cfg := &DatabaseConfig{
		Type:        string(tp),
		ConnInfo:    string(connInfo),
		AutoMigrate: migrate,
	}
	return cfg, nil
}

func (s *ZooKeeperServer) getApiConfig(basePath string, conn *zk.Conn) (*ApiConfig, error) {
	apiAddr, _, err := conn.Get(basePath + PropertyApiAddr)
	if err != nil {
		return nil, fmt.Errorf("get %s error: %v", PropertyApiAddr, err)
	}
	if string(apiAddr) == "" {
		return nil, ErrInvalidApiAddr
	}

	apiCfg := &ApiConfig{
		APIAddr: string(apiAddr),
	}
	return apiCfg, nil
}

func (s *ZooKeeperServer) getGlobalConfig(basePath string, conn *zk.Conn) (*GlobalConfig, error) {
	secretKey, err := s.getSecretKey(basePath, conn)
	if err != nil {
		return nil, err
	}

	logLevel, _, err := conn.Get(basePath + PropertyLoglevel)
	if err != nil {
		logLevel = []byte("")
	}
	logEncoder, _, err := conn.Get(basePath + PropertyLogEncoder)
	if err != nil {
		logEncoder = []byte("")
	}

	globalCfg := &GlobalConfig{
		SecretKey:  secretKey,
		LogLevel:   GetLogLevel(string(logLevel)),
		LogEncoder: GetLogEncoder(string(logEncoder)),
	}
	return globalCfg, nil
}

func (s *ZooKeeperServer) getSecretKey(basePath string, conn *zk.Conn) (string, error) {
	secretKey, _, err := conn.Get(basePath + PropertySecretKey)
	if err != nil {
		return "", fmt.Errorf("get %s error: %v", PropertySecretKey, err)
	}
	if string(secretKey) == "" {
		return "", ErrInvalidSecretKey
	}
	return string(secretKey), nil
}

func (s *ZooKeeperServer) getApiAddr(basePath string, conn *zk.Conn) (string, error) {
	apiAddr, _, err := conn.Get(basePath + PropertyApiAddr)
	if err != nil {
		return "", fmt.Errorf("get %s error: %v", PropertyApiAddr, err)
	}
	if string(apiAddr) == "" {
		return "", ErrInvalidApiAddr
	}
	return string(apiAddr), nil
}
