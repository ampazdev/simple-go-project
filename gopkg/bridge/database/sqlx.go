package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type sqlxDB struct {
	DB *sqlx.DB
}

// Config database configuration
type Config struct {
	Driver string     `json:"driver" yaml:"driver"`
	Master DBParams   `json:"master" yaml:"master"`
	Slave  []DBParams `json:"slave" yaml:"slave"`
}

type DBParams struct {
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"db_name" yaml:"db_name"`
	Address  string `json:"address" yaml:"address"`
	Port     int    `json:"port" yaml:"port"`
	SSLMode  string `json:"ssl_mode" yaml:"ssl_mode"`
}

// Constructor of sqlxDB
func NewSqlxDB() *sqlxDB {
	return &sqlxDB{}
}

// TODO: implement slave in database
func (s *sqlxDB) Connect(_ context.Context, cfg Config) (*sqlxDB, error) {
	dbMasterURI := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.Master.User,
		cfg.Master.Password,
		cfg.Master.DBName,
		cfg.Master.Address,
		cfg.Master.Port,
		cfg.Master.SSLMode,
	)

	db, err := sqlx.Connect(cfg.Driver, dbMasterURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &sqlxDB{
		DB: db,
	}, nil
}
