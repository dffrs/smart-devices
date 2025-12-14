package internal

import (
	"os"

	sh "smart-devices/internal/shelly"
)

const timeout = 10 // seconds

func CreateOutlet() (*sh.Outlet, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	host := os.Getenv("HOST")
	outlet := sh.NewOutlet(host, timeout)

	return outlet, nil
}
