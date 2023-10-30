package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(
	db *sqlx.DB,
) *Client {
	return &Client{
		db: db,
	}
}

func (c *Client) Exec(
	ctx context.Context,
	query string,
	arg any,
) (sql.Result, error) {
	namedQuery, namedArgs, err := c.prepare(query, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to run c.prepare: %w", err)
	}

	result, err := c.db.ExecContext(ctx, namedQuery, namedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to run c.db.ExecContext %w", err)
	}

	return result, nil
}

func (c *Client) Query(
	ctx context.Context,
	query string,
	arg interface{},
) (*sqlx.Rows, error) {
	namedQuery, namedArgs, err := c.prepare(query, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to run c.prepare: %w", err)
	}

	rows, err := c.db.QueryxContext(ctx, namedQuery, namedArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to run c.db.QueryxContext: %w", err)
	}

	return rows, nil
}

func (c *Client) prepare(
	query string,
	arg any,
) (string, []any, error) {
	namedQuery, namedArgs, err := c.db.BindNamed(query, arg)
	if err != nil {
		return "", nil, fmt.Errorf("failed to run c.db.BindNamed: %w", err)
	}

	namedQuery, namedArgs, err = sqlx.In(namedQuery, namedArgs...)
	if err != nil {
		return "", nil, fmt.Errorf("failed to run sqlx.In: %w", err)
	}

	return namedQuery, namedArgs, nil
}
