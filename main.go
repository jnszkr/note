package main

import (
	"flag"
	"os"
)

func main() {

	var s string

	flag.StringVar(&s, "s", "", "Search in notes.")

	flag.Parse()

	switch {
	case len(s) > 0:
		search(s)
	default:
		add(os.Args[1:])
	}
}
