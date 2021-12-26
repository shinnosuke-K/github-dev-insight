package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	envFile = ".env"
)

var (
	app = &AppEnv{}
)

func Get() *AppEnv {
	return app
}

func Load() error {
	if err := godotenv.Load(filepath.Join(os.Getenv("PWD"), envFile)); err != nil {
		return fmt.Errorf("failed to load .env. %w", err)
	}
	if err := envconfig.Process("", app); err != nil {
		return fmt.Errorf("failed to mapping env. %w", err)
	}
	return nil
}
