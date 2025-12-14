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

func Run(outlet *shelly.Outlet) error {
	state, status := LoadFlags()

	switch true {

	// deal with get status/methods options
	case *status != "":
		switch *status {
		case getStatus:
			output := outlet.GetStatus()
			PrettyPrint(output)
		case getMethods:
			output := outlet.ListMethods()
			PrettyPrint(output)
		default:
			return errors.New("unrecognized option for get option")
		}

	// deal with turn on/off option
	case *state != "":
		var isOn bool
		switch *state {
		case turnOn:
			isOn = true
		case turnOff:
			isOn = false
		default:
			return errors.New("unrecognized option for turn option")
		}

		outlet.Turn(isOn)
	}

	return nil
}
