package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jnszkr/note/searcher"
)

func main() {

	var s string

	flag.StringVar(&s, "s", "", "Search in notes.")

	flag.Parse()

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case len(s) > 0:
		sr := searcher.New(currDir, os.Stdout)
		sr.Search(s)
	case len(os.Args) > 1:
		add(os.Args[1:])
	default:
		path := filepath.Join(currDir, ".notes")
		fmt.Println(display(path))
	}
}
