package bridge

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	sqlx.QueryerContext
	sqlx.ExecerContext
}
