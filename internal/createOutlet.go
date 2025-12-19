package internal

import (
	sh "smart-devices/internal/shelly"
)

const timeout = 10 // seconds

func CreateOutlet() (*sh.Outlet, error) {
	conf, err := loadEnv()
	if err != nil {
		return nil, err
	}

	outlet := sh.NewOutlet(conf.host, timeout)

	return outlet, nil
}
