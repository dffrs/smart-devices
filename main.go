package main

import (
	"flag"
	"log"
	"os"

	outlet "smart-devices/internal/shelly/outlet"

	"github.com/joho/godotenv"
)

const (
	timeout = 10 // seconds
	turnOn  = "on"
	turnOff = "off"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file\n%v", err)
	}

	state := flag.String("turn", turnOn, "Turn outlet on & off")
	flag.Parse()

	host := os.Getenv("HOST")

	outlet := outlet.NewShelly(host, timeout)

	isOn := *state != turnOff

	outlet.Turn(isOn)
}
