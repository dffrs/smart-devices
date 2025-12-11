package internal

import (
	"encoding/json"
	"fmt"
)

type Shelly interface {
	NewShelly() *Shelly
	GetStatus() any
	ListMethods() any
}

func PrettyPrint(arg any) {
	b, _ := json.MarshalIndent(arg, "", "  ")
	fmt.Println(string(b))
}
