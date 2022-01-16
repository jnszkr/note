package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jnszkr/note/internal/formatter"
	"github.com/jnszkr/note/internal/reader"
	"github.com/jnszkr/note/internal/searcher"
	"github.com/jnszkr/note/internal/writer"
)

func main() {

	var s string
	var r bool

	flag.StringVar(&s, "s", "", "Search in notes")
	flag.BoolVar(&r, "r", false, "Search recursively")

	flag.Parse()

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case len(s) > 0:
		// search parameter provided
		sr := searcher.New(currDir, os.Stdout)
		sr.Search(s, r)
	case len(os.Args) > 1:
		w, err := writer.New(filepath.Join(currDir, searcher.FileName))
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()
		err = w.WriteNote(strings.Join(os.Args[1:], " "))
		if err != nil {
			log.Fatal(err)
		}
	default:
		// no arguments provided
		path := filepath.Join(currDir, searcher.FileName)
		r, err := reader.New(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(formatter.Format(r))
	}
}
