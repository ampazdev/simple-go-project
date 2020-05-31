package config

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

type Config struct {
	AppName string   `yaml:"appname"`
	DB      Database `yaml:"db"`
	Redis   Redis    `yaml:"redis"`
	Cache   redis.Conn
	Rest
}

// NewConfig creates new Config by reading values stored in config file
func NewConfig(filepath string) (*Config, error) {
	cfg := Config{}

	cfgReader := viper.New()
	cfgReader.SetConfigFile(filepath)
	err := cfgReader.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = cfgReader.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
