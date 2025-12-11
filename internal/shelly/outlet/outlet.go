package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	in "smart-devices/internal/models"
)

type Outlet struct {
	host   string
	client *http.Client
}

func NewShelly(host string, timeoutInSeconds int) *Outlet {
	client := &http.Client{Timeout: time.Duration(timeoutInSeconds) * time.Second}
	return &Outlet{host: host, client: client}
}

func (o *Outlet) GetStatus() *in.GetStatus {
	endpoint := "/rpc/Shelly.GetStatus"

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		log.Fatalf("Failed to call GetStatus endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var status *in.GetStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		log.Fatalf("Failed to decode JSON in GetStatus\n%v", err)
	}
	return status
}

func (o *Outlet) ListMethods() *in.ListMethods {
	endpoint := "/rpc/Shelly.ListMethods"

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		log.Fatalf("Failed to call ListMethods endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var methods *in.ListMethods
	err = json.NewDecoder(resp.Body).Decode(&methods)
	if err != nil {
		log.Fatalf("Failed to decode JSON in ListMethods\n%v", err)
	}
	return methods
}
