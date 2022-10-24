package config

import (
	"github.com/jinzhu/configor"
	"log"
)

const _DefaultConfigFileName = "config.toml"

var _config *Config

func GetConfig() *Config {
	if _config == nil {
		_config = &Config{}
		err := configor.Load(_config, _DefaultConfigFileName)
		if err != nil {
			log.Panic("未找到配置文件:", err)
		}
	}
	return _config
}

type Config struct {
	DB    map[string]DBConfig
	Redis *struct {
		Addr     string
		Password string
		DB       int
	}
	Server struct {
		Debug   bool
		ShowSql bool
		Port    int `default:"3000"`
	}
	AppConfig map[string]string
}

type DBConfig struct {
	Type string `default:"mysql"`
	Url  string
}

func (c *Config) GetAppCfg(name string) string {
	return c.AppConfig[name]
}

const DefaultDBName = "default"

func (c *Config) GetDefaultDB() DBConfig {
	return c.DB[DefaultDBName]
}

func (c *Config) GetDB(name string) DBConfig {
	return c.DB[name]
}
