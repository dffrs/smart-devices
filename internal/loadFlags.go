package internal

import "flag"

func LoadFlags() (*string, *string) {
	status := flag.String("get", "", "Get Status report from outlet")
	state := flag.String("turn", "", "Turn outlet on & off")
	flag.Parse()

	return state, status
}
