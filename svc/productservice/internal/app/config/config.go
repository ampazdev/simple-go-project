package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName string   `yaml:"appname"`
	DB      Database `yaml:"db"`
	Rest
}

// NewConfig creates new Config by reading values stored in config file
func NewConfig(filepath string) (Config, error) {
	cfg := Config{}

	cfgReader := viper.New()
	cfgReader.SetConfigFile(filepath)
	err := cfgReader.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = cfgReader.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
