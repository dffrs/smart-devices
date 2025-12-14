package main

import (
	"log"

	in "smart-devices/internal"
)

func main() {
	outlet, err := in.CreateOutlet()
	if err != nil {
		log.Fatal(err)
	}

	err = in.Run(outlet)
	if err != nil {
		log.Fatal(err)
	}
}
