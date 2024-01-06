package app

import (
	"fmt"

	"layout/internal/app/config"
	"layout/internal/app/transport/api"
)

func Run(cfg *config.API) error {
	api := api.New(&cfg.HTTP)
	if err := api.Listen(fmt.Sprintf("%s:%d", cfg.HTTP.ListenHost, cfg.HTTP.ListenPort)); err != nil {
		return err
	}

	return nil
}
