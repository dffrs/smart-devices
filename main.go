package main

import (
	"flag"
	"log"
	"os"

	internal "smart-devices/internal/shelly"
	outlet "smart-devices/internal/shelly/outlet"

	"github.com/joho/godotenv"
)

const (
	timeout = 10 // seconds
	turnOn  = "on"
	turnOff = "off"

	getStatus  = "status"
	getMethods = "methods"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file\n%v", err)
	}

	state := flag.String("turn", "", "Turn outlet on & off")
	status := flag.String("get", "", "Get Status report from outlet")
	flag.Parse()

	host := os.Getenv("HOST")
	outlet := outlet.NewShelly(host, timeout)

	// TODO: clean me up
	switch true {

	// deal with get status/methods options
	case *status != "":
		switch *status {
		case getStatus:
			output := outlet.GetStatus()
			internal.PrettyPrint(output)
		case getMethods:
			output := outlet.ListMethods()
			internal.PrettyPrint(output)
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
