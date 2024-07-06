package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

//TODO:hot update

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
	MaxAge      int    `yaml:"maxAge"`
	MaxBackups  int    `yaml:"maxBackups"`
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
var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		b, err := ioutil.ReadFile(_CONFIG_FILE_PATH)
		if err != nil {
			panic(err)
		}

		serverConf = &ServerConfig{}
		err = yaml.Unmarshal(b, serverConf)
		if err != nil || !serverConf.check() {
			panic(err)
		}
		log.Printf("load config succes. %v", serverConf)
	})
}

func GetServerConfig() *ServerConfig {
	return serverConf
}

func (c *ServerConfig) check() bool {
	if c.LogConfig == nil {
		return false
	}
	if c.RedisConfig == nil {
		return false
	}
	if c.MysqlConfig == nil {
		return false
	}
	return true
}
