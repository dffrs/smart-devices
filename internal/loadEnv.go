package internal

import (
	"errors"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("failed to load .env file\n%v")
	}
	return nil
}
