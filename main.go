package main

import (
	"log"
	"os"

	shelly "smart-devices/internal/shelly"
	outlet "smart-devices/internal/shelly/outlet"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	host := os.Getenv("HOST")

	outlet := outlet.NewShelly(host, 3)

	status := outlet.GetStatus()

	shelly.PrettyPrint(&status)
}
