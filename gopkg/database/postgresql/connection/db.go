package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	Database struct {
		*sqlx.DB
	}
)

func New(config *Config) (*Database, error) {
	db, err := sqlx.Connect("postgres", config.DatabaseURI)
	if err != nil {
		return nil, fmt.Errorf("error connect to DB")
	}
	return &Database{db}, nil
}