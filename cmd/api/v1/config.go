package v1

import (
	"fmt"
	"sync"

	envx "go.strv.io/env"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

const dotenvPath = "v1/.env"

var (
	once sync.Once

	validate = validator.New()
)

type Config struct {
	Port        int    `env:"PORT" validate:"required"`
	DatabaseURL string `env:"DATABASE_URL" validate:"required"`
}

func LoadConfig() (Config, error) {
	loaddotenv(dotenvPath)

	cfg := Config{}
	if err := envx.Apply(&cfg); err != nil {
		return cfg, err
	}
	fmt.Println(cfg)
	if err := validate.Struct(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil

}

func MustLoadConfig() Config {
	cfg, err := LoadConfig()
	fmt.Println(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func loaddotenv(path string) {
	once.Do(func() {
		if path == "" {
			path = ".env"
		}

		_ = godotenv.Load(dotenvPath)
		_ = godotenv.Load(dotenvPath + ".common")
	})
}
