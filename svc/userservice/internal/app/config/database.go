package config

import (
	_ "github.com/lib/pq" // to import postgres driver
)

type Database struct {
	Engine string `yaml:"engine"`
	Master Params `yaml:"master"`
	Slave  Params `yaml:"slave"`
}

type Params struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	Port     int    `yaml:"port"`
}
