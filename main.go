package main

import (
	"flag"
	"log"
	"os"

	"github.com/jnszkr/note/searcher"
)

func main() {

	var s string

	flag.StringVar(&s, "s", "", "Search in notes.")

	flag.Parse()

	switch {
	case len(s) > 0:
		currDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		sr := searcher.New(currDir, os.Stdout)
		sr.Search(s)
	default:
		add(os.Args[1:])
	}
}
