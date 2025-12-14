package internal

import (
	"fmt"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file\n%v", err.Error())
	}
	return nil
}
