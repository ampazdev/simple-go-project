package config

import (
	"github.com/ampazdev/simple-go-project/gopkg/bridge/database"
	_ "github.com/lib/pq" // to import postgres driver
)

type Database struct {
	SqlxConfig database.Config `yaml:"sqlx_config"`
}
