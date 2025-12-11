package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	endpoint := "/rpc/Switch.GetStatus?id=0"

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://%s%s", host, endpoint))
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()

	var temp Result
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		log.Fatalf("%v", err)
	}

	b, _ := json.MarshalIndent(temp, "", "  ")
	fmt.Println(string(b))
}
