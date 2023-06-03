package opendbgo

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(ctx context.Context, driver string, dsn string, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	if err = initTable(ctx, db, path); err != nil {
		return nil, err
	}

	return db, nil
}

func initTable(ctx context.Context, db *sql.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if _, err = db.ExecContext(ctx, string(data)); err != nil {
		return err
	}
	return nil
}
