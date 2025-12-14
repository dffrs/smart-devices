package internal

import (
	"errors"

	"smart-devices/internal/shelly"
)

const (
	turnOn  = "on"
	turnOff = "off"

	getStatus  = "status"
	getMethods = "methods"
)

func dealWithStatus(status *string, outlet *shelly.Outlet) error {
	var output any
	var err error

	switch *status {
	case getStatus:
		output, err = outlet.GetStatus()
	case getMethods:
		output, err = outlet.ListMethods()
	default:
		err = errors.New("unrecognized option for get option")
	}

	if err != nil {
		return err
	}

	PrettyPrint(output)

	return nil
}

func getTurnOption(state *string) (bool, error) {
	isOn := false

	switch *state {
	case turnOn:
		isOn = true
	case turnOff:
		isOn = false
	default:
		return false, errors.New("unrecognized option for turn option")
	}

	return isOn, nil
}

func dealWithState(state *string, outlet *shelly.Outlet) error {
	isOn, err := getTurnOption(state)
	if err != nil {
		return err
	}

	_, err = outlet.Turn(isOn)
	if err != nil {
		return err
	}

	return nil
}

func Run(outlet *shelly.Outlet) error {
	state, status := loadFlags()

	switch true {

	// deal with get status/methods options
	case *status != "":
		err := dealWithStatus(status, outlet)
		if err != nil {
			return err
		}

	// deal with turn on/off option
	case *state != "":
		err := dealWithState(state, outlet)
		if err != nil {
			return err
		}
	}

	return nil
}
