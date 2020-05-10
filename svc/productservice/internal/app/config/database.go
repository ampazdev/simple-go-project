package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // to import postgres driver
)

type Database struct {
	Engine string `yaml:"engine"`
	Master Params `yaml:"master"`
	Slave  Params `yaml:"slave"`
}

type Params struct {
	Address  string `yaml:"address"`
	User     string `yaml:"userservices"`
	Password string `yaml:"password"`
	DbName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
	Port     int    `yaml:"port"`
}

func NewDatabase(cfg Config) (*sqlx.DB, error) {
	databaseURI := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.DB.Master.User,
		cfg.DB.Master.Password,
		cfg.DB.Master.DbName,
		cfg.DB.Master.Address,
		cfg.DB.Master.Port,
		cfg.DB.Master.SSLMode,
	)

	db, err := sqlx.Connect("postgres", databaseURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
