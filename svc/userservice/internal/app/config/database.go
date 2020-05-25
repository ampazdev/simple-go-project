package config

import (
	_ "github.com/lib/pq" // to import postgres driver
)

type Database struct {
	Engine string `yaml:"engine"`
	Master Params `yaml:"master"`
	Slave  Params `yaml:"slave"`
}

type Redis struct {
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
}

type Params struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Port     int    `yaml:"port"`
}
