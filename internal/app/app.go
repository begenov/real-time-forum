package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/begenov/real-time-forum/internal/config"
	delivery "github.com/begenov/real-time-forum/internal/delivery/http"
	"github.com/begenov/real-time-forum/internal/repository"
	"github.com/begenov/real-time-forum/internal/server"
	"github.com/begenov/real-time-forum/internal/service"
	"github.com/begenov/real-time-forum/pkg/auth"
	"github.com/begenov/real-time-forum/pkg/hash"
	opendb "github.com/begenov/real-time-forum/pkg/open_db"
)

const path = "./migration/init.up.sql"

func Run(cfg *config.Config) error {
	db, err := opendb.OpenDB(context.Background(), cfg.Database.Driver, cfg.Database.Dsn, path)
	if err != nil {
		return err
	}

	hash := hash.NewHash(cfg.Hash.Cost)

	manager := auth.NewManager()

	repo := repository.NewRepository(db)

	service := service.NewService(repo, hash, manager, cfg)

	handler := delivery.NewHandler(service)

	srv := server.NewServer(cfg, handler.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err)
		}
	}()

	log.Println("Server started", cfg.Server.Port)

	quit := make(chan os.Signal, 1)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}

	return nil
}
