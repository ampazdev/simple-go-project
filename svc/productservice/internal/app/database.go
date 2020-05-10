package app

import (
	"database/sql"
	"fmt"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app/config"
)

// TODO: move to bridge
type Database struct {
	*sql.DB
}

func newDatabase(cfg *config.Database) (*Database, error) {
	connection := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.Master.Address, cfg.Master.Port, cfg.Master.User, cfg.Master.Password, cfg.Master.DbName, cfg.Master.SSLMode)

	//val := url.Values{}
	//val.Add("parseTime", "1")
	//val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s", connection)
	dbConn, err := sql.Open(cfg.Engine, dsn)
	if err != nil {
		return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{
		dbConn,
	}, nil
}
