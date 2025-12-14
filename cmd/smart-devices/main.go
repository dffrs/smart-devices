package main

import (
	"flag"
	"log"

	in "smart-devices/internal"
)

const (
	turnOn  = "on"
	turnOff = "off"

	getStatus  = "status"
	getMethods = "methods"
)

func main() {
	state := flag.String("turn", "", "Turn outlet on & off")
	status := flag.String("get", "", "Get Status report from outlet")
	flag.Parse()

	outlet, err := in.CreateOutlet()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: clean me up
	switch true {

	// deal with get status/methods options
	case *status != "":
		switch *status {
		case getStatus:
			output := outlet.GetStatus()
			in.PrettyPrint(output)
		case getMethods:
			output := outlet.ListMethods()
			in.PrettyPrint(output)
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
			log.Fatalf("Unrecognized option for turn option\n")
		}

		outlet.Turn(isOn)
	}
}
