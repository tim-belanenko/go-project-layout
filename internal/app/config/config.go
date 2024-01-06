package config

import "github.com/ilyakaznacheev/cleanenv"

type HTTP struct {
	ListenPort int    `env:"LISTEN_PORT" env-default:"8080" validate:"gt=1023,lte=65536"`
	ListenHost string `env:"LISTEN_HOST" env-default:"localhost"`
}

type DSN struct {
	POSTGRES string `env:"POSTGRES" env-required:"true" validate:"url"`
}

type API struct {
	HTTP HTTP `env-prefix:"HTTP_"`
	DSN  DSN  `env-prefix:"DSN_"`
}

type Migration struct {
	DSN DSN `env-prefix:"DSN_"`
}

func ReadEnv(cfg any) error {
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return err
	}

	return nil
}
