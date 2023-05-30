package opendbgo

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(ctx context.Context, driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
