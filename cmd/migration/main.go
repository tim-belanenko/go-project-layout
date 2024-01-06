package main

import (
	"layout/internal/app/config"
	"layout/internal/pkg/ctxlogger"
	"layout/internal/pkg/migration"
	"layout/internal/pkg/validator"
)

func main() {
	log := ctxlogger.DefaultLogger()

	cfg := &config.Migration{}

	if err := config.ReadEnv(cfg); err != nil {
		log.Fatal(err.Error())
	}

	if err := validator.Validate().Struct(cfg); err != nil {
		log.Fatal(err.Error())
	}

	log.Debug("start migration")
	migration, err := migration.New(cfg.DSN.POSTGRES)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := migration.Up(); err != nil {
		log.Fatal(err.Error())
	}
}
