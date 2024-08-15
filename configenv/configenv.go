package configenv

import (
	"fmt"
	"os"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// LoadConfig function used to load environment variab from env variable
func LoadConfig(appName string) (*dstructures.EnvConfig, error) {

	var cfg dstructures.EnvConfig

	if _, err := os.Stat(".env"); err == nil {
		println("[ENV] Load env variables from .env")
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}

	}

	err := envconfig.Process(appName, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
