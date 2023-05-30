package app

import (
	"context"

	"github.com/begenov/real-time-forum/internal/config"
	"github.com/begenov/real-time-forum/internal/repository"
	opendb "github.com/begenov/real-time-forum/pkg/open_db"
)

const path = "./migration/init.up.sql"

func Run(cfg *config.Config) error {
	db, err := opendb.OpenDB(context.Background(), cfg.Database.Driver, cfg.Database.Dsn, path)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db)

	return nil
}
