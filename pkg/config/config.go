package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/toolkits/file"
)

// GlobalConfig global config
var GlobalConfig *Config

const (
	// GATHER_ENV 读取的环境变量
	GATHER_ENV = "PLAN_ENV"
	// DEFAULT_ENV 默认环境变量
	DEFAULT_ENV = "development"
	// CONFIG_DIR 配置文件目录
	CONFIG_DIR         = "config"
	CONFIG_FILE_SUFFIX = ".json"
)

// MySQLConfig mysql config
type MySQLConfig struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"db_name,omitempty"`
}

type RedisConfig struct {
	Addr       string `json:"addr,omitempty"`
	Password   string `json:"password,omitempty"`
	DB         int    `json:"db,omitempty"`
	MaxRetries int    `json:"max_retries"`
}

type ESConfig struct {
	Endpoints []string `json:"endpoints"`
}

type KafkaConfig struct {
	Endpoints []string `json:"endpoints"`
}

// Config common config
type Config struct {
	Mysql *MySQLConfig `json:"mysql,omitempty"`
	Redis *RedisConfig `json:"redis,omitempty"`
	ES    *ESConfig    `json:"es,omitempty"`
	Kafka *KafkaConfig `json:"kafka,omitempty"`
	Debug bool         `json:"debug,omitempty"`
	Env   string       `json:"-"`
}

// LoadConfig load config
func LoadConfig(configPath string) {
	if configPath == "" {
		configPath = CONFIG_DIR
	}
	var err error
	env := os.Getenv(GATHER_ENV)
	if env == "" {
		env = DEFAULT_ENV
	}
	fileName := env + CONFIG_FILE_SUFFIX
	filePath := fmt.Sprintf("%s/%s", configPath, fileName)
	var configContent string
	configContent, err = file.ToTrimString(filePath)
	if err != nil {
		err = fmt.Errorf("read config file: %v, err: %v", fileName, err)
	}
	var c Config
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		err = fmt.Errorf("arse config file: %v, err: %v", fileName, err)
	}
	c.Env = env

	GlobalConfig = &c

	if err != nil {
		panic(err)
	}
	return
}
