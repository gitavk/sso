package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string        `env:"ENV" env-required:"true"`
	StoragePath string        `env:"STORAGE_PATH" env-required:"true"`
	TokenTTL    time.Duration `env:"TOKEN_TTL" env-default:"1h"`
	GRPC        GRPCConfig    `env-prefix:"GRPC_"`
}

type GRPCConfig struct {
	Port    int           `env:"PORT"`
	Timeout time.Duration `env:"TIMEOUT"`
}

func MustLoad() *Config {
	// Load .env file first (if it exists)
	_ = godotenv.Load(".env")

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Println(err)
		panic("failed to read config.")
	}
	return &cfg

}
