package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type accessor interface {
	Exec(context.Context, string, any) (sql.Result, error)
	Query(context.Context, string, any) (*sqlx.Rows, error)
}
