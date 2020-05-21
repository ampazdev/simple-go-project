package app

import (
	"context"

	"github.com/ampazdev/simple-go-project/gopkg/bridge"
	"github.com/ampazdev/simple-go-project/gopkg/bridge/database"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app/config"
)

type Bridges struct {
	DB bridge.Database
}

func newBridges(cfg *config.Database) (*Bridges, error) {
	// Database init
	db, err := database.NewSqlxDB().Connect(context.TODO(), cfg.SqlxConfig)
	if err != nil {
		return nil, err
	}

	return &Bridges{
		DB: db,
	}, nil
}

func (b *Bridges) Close() []error {
	var errs []error

	// TODO: add db conn close
	//errs = append(errs, b.DB.Close())

	return errs
}
