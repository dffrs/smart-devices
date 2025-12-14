// Package shelly provides low-level abstractions for interacting with IoT outlet device
//
// This package contains an Outlet type, which represents a network-connected
// smart power outlet and exposes a small, focused API for querying device status,
// listing supported RPC methods, and toggling power state.
package shelly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	tp "smart-devices/api/types"
)

type Outlet struct {
	host   string
	client *http.Client
}

func NewOutlet(host string, timeoutInSeconds int) *Outlet {
	client := &http.Client{Timeout: time.Duration(timeoutInSeconds) * time.Second}
	return &Outlet{host: host, client: client}
}

func (o *Outlet) GetStatus() (*tp.GetStatus, error) {
	endpoint := "/rpc/Shelly.GetStatus"

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to call GetStatus endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var status *tp.GetStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON in GetStatus\n%v", err)
	}
	return status, nil
}

func (o *Outlet) ListMethods() (*tp.ListMethods, error) {
	endpoint := "/rpc/Shelly.ListMethods"

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to call ListMethods endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var methods *tp.ListMethods
	err = json.NewDecoder(resp.Body).Decode(&methods)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON in ListMethods\n%v", err)
	}
	return methods, nil
}

func (o *Outlet) Turn(onOrOff bool) (*tp.Turn, error) {
	endpoint := fmt.Sprintf("/rpc/Switch.Set?id=0&on=%t", onOrOff)

	resp, err := o.client.Get(fmt.Sprintf("http://%s%s", o.host, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to call ListMethods endpoint\n%v", err)
	}
	defer resp.Body.Close()

	var temp *tp.Turn
	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON in ListMethods\n%v", err)
	}
	return temp, nil
}
