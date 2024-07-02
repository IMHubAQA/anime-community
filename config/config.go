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
	LogConfig   *LogConfig   `yaml:"logConfig"`
	MysqlConfig *MysqlConfig `yaml:"mysqlConfig"`
	RedisConfig *RedisConfig `yaml:"redisConfig"`
}

type LogConfig struct {
	FilePath    string `yaml:"filePath"`
	ErrFilePath string `yaml:"errFilePath"`
	MaxSize     int    `yaml:"maxSize"`
	MaxDay      int    `yaml:"maxDay"`
}
type MysqlConfig struct {
	Protocol string `yaml:"protocol"`
	Addr     string `yaml:"addr"`
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
	DbName   string `yaml:"dbName"`
	Charset  string `yaml:"charset"`
}
type RedisConfig struct {
	Addr     []string `yaml:"addr"`
	PassWord string   `yaml:"passWord"`
}

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
