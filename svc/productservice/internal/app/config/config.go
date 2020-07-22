package config

import (
	"github.com/ampazdev/simple-go-project/svc/productservice/constant"
	"go.uber.org/config"
)

type Config struct {
	AppName string   `yaml:"app_name"`
	DB      Database `yaml:"db"`
	Rest    Rest     `yaml:"rest"`
	GRPC    GRPC     `yaml:"grpc"`
}

// NewConfig creates new Config by reading values stored in config file
func NewConfig(filepath string) (*Config, error) {
	cfg := Config{}

	provider, err := config.NewYAML(config.File(constant.ConfigProjectFilepath))
	if err != nil {
		return nil, err
	}

	err = provider.Get("").Populate(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
