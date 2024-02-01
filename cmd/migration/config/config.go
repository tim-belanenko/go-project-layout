package config

import "github.com/ilyakaznacheev/cleanenv"

type DSN struct {
	POSTGRES string `env:"POSTGRES" env-required:"true" validate:"url"`
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
