package app

import (
	"context"
	"fmt"
	"log"

	"layout/cmd/api/config"
	"layout/internal/app/repository/postgres"
	"layout/internal/app/transport/api"
	"layout/internal/app/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(cfg *config.API) error {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DSN.POSTGRES)
	if err != nil {
		return err
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repo := postgres.New(pool)
	uc := usecase.New(repo)

	api := api.New(&cfg.HTTP, uc)
	if err := api.Listen(fmt.Sprintf("%s:%d", cfg.HTTP.ListenHost, cfg.HTTP.ListenPort)); err != nil {
		return err
	}

	return nil
}
