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
	var error error

	// TODO: getStatus and listMethods should return errors
	switch *status {
	case getStatus:
		output = outlet.GetStatus()
	case getMethods:
		output = outlet.ListMethods()
	default:
		error = errors.New("unrecognized option for get option")
	}

	if error != nil {
		return error
	}

	PrettyPrint(output)

	return nil
}

func dealWithState(state *string, outlet *shelly.Outlet) error {
	var isOn bool
	var err error

	switch *state {
	case turnOn:
		isOn = true
	case turnOff:
		isOn = false
	default:
		err = errors.New("unrecognized option for turn option")
	}

	if err != nil {
		return err
	}

	outlet.Turn(isOn)

	return nil
}

func Run(outlet *shelly.Outlet) error {
	state, status := LoadFlags()

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
