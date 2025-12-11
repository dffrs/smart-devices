package main

import (
	"log"
	"os"

	shelly "smart-devices/internal/shelly"
	outlet "smart-devices/internal/shelly/outlet"

	"github.com/joho/godotenv"
)

type Result struct {
	ID          int     `json:"id"`
	Source      string  `json:"source"`
	Output      bool    `json:"output"`
	APower      float64 `json:"apower"`
	Voltage     float64 `json:"voltage"`
	Freq        float64 `json:"freq"`
	Current     float64 `json:"current"`
	AEnergy     Energy  `json:"aenergy"`
	RetAEnergy  Energy  `json:"ret_aenergy"`
	Temperature Temp    `json:"temperature"`
}

type Energy struct {
	Total    float64   `json:"total"`
	ByMinute []float64 `json:"by_minute"`
	MinuteTS int64     `json:"minute_ts"`
}

type Temp struct {
	TC float64 `json:"tC"`
	TF float64 `json:"tF"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
	}

	host := os.Getenv("HOST")

	outlet := outlet.NewShelly(host, 3)

	temp := outlet.GetStatus()

	shelly.PrettyPrint(&temp)
}
