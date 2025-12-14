// Package internal contains shared internal helpers and abstractions used by the application.
package internal

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(arg any) {
	b, _ := json.MarshalIndent(arg, "", "  ")
	fmt.Println(string(b))
}
