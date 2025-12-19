package internal

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	host string
}

const envFile = "sd.env"

func loadEnv() (*config, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	envPath := fmt.Sprintf("%s/%s", userConfigDir, envFile)

	err = godotenv.Load(envPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s file. Create it, if it does not exist, with content HOST = <IP of outlet>", envPath)
	}

	return &config{host: os.Getenv("HOST")}, nil
}
