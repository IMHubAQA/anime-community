package config

import (
	"io/ioutil"

	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/yaml.v3"
)

const (
	_CONFIG_FILE_PATH = "./conf/config.yaml"
)

type ServerConfig struct {
	LogConfig   interface{} `yaml:"logConfig"`
	MysqlConfig interface{} `yaml:"mysqlConfig"`
	RedisConfig interface{} `yaml:"redisConfig"`
}

type LogConfig struct{}
type MysqlConfig struct{}
type RedisConfig struct{}

var serverConf *ServerConfig

func init() {
	b, err := ioutil.ReadFile(_CONFIG_FILE_PATH)
	if err != nil {
		panic(err)
	}

	serverConf = &ServerConfig{}
	err = yaml.Unmarshal(b, serverConf)
	if err != nil {
		panic(err)
	}

	logs.Info("load config succes. %v", serverConf)
}

func GetServerConfig() *ServerConfig {
	return serverConf
}
