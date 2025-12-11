package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Outlet struct {
	host   string
	client *http.Client
}

func NewShelly(host string, timeoutInSeconds int) *Outlet {
	client := &http.Client{Timeout: time.Duration(timeoutInSeconds) * time.Second}
	return &Outlet{host: host, client: client}
}

func (o *Outlet) GetStatus() any {
	endpoint := "/rpc/Shelly.GetStatus"

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		log.Fatalf("Failed to call GetStatus endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var temp any
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		log.Fatalf("Failed to decode JSON in GetStatus\n%v", err)
	}
	return temp
}

func (o *Outlet) ListMethods() any {
	return nil
}
