package main

import (
	"log"
	"os"

	shelly "smart-devices/internal/shelly"
	outlet "smart-devices/internal/shelly/outlet"

	"github.com/joho/godotenv"
)

const timeout = 10 // seconds

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	host := os.Getenv("HOST")

	outlet := outlet.NewShelly(host, timeout)

	// status := outlet.GetStatus()
	methods := outlet.ListMethods()

	// shelly.PrettyPrint(status)
	shelly.PrettyPrint(methods)
}
