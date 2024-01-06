package main

import (
	"layout/internal/app"
	"layout/internal/app/config"
	"layout/internal/pkg/ctxlogger"
	"layout/internal/pkg/validator"
)

func main() {
	log := ctxlogger.DefaultLogger()

	cfg := &config.API{}

	if err := config.ReadEnv(cfg); err != nil {
		log.Fatal(err.Error())
	}

	if err := validator.Validate().Struct(cfg); err != nil {
		log.Fatal(err.Error())
	}

	if err := app.Run(cfg); err != nil {
		log.Fatal(err.Error())
	}
}
