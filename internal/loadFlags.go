package internal

import "flag"

func loadFlags() (*string, *string) {
	status := flag.String("get", "", "Get Status report from outlet")
	state := flag.String("set", "", "Turn outlet on & off")
	flag.Parse()

	return state, status
}
