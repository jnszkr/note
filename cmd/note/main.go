package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jnszkr/note/internal/files"
	"github.com/jnszkr/note/internal/notes"
	"github.com/jnszkr/note/internal/writer"
)

const fileName = ".notes"

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
		search(currDir, r, s)
	case len(os.Args) > 1:
		write(currDir)
	default:
		show(currDir)
	}
}

func show(currDir string) {
	format := notes.Formatter()
	fmt.Println(format(notes.Reader(filepath.Join(currDir, fileName))))
}

func write(currDir string) {
	w, err := writer.New(filepath.Join(currDir, fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	err = w.WriteNote(strings.Join(os.Args[1:], " "))
	if err != nil {
		log.Fatal(err)
	}
}

func search(currDir string, r bool, s string) {
	t0 := time.Now()
	defer func() {
		if os.Getenv("DEBUG") != "" {
			fmt.Printf("%sSearch duration: %v%s\n", notes.YellowColor, time.Since(t0), notes.ResetColor)
		}
	}()

	fs := files.Find(currDir, ".notes", r)

	var wg sync.WaitGroup
	format := notes.Formatter(notes.WithPrefix("    "), notes.WithHighlight(s, notes.RedColor))
	for f := range fs {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()

			t1 := time.Now()

			var b bytes.Buffer

			res := format(
				notes.Filter(
					notes.Reader(f),
					func(note *notes.Note) bool {
						return strings.Contains(note.Text, s)
					},
				),
			)

			b.WriteString(f)
			if os.Getenv("DEBUG") != "" {
				b.WriteString("   ")
				b.WriteString(notes.Highlight("", notes.YellowColor, time.Since(t1).String()))
			}
			b.WriteString("\n")
			b.WriteString(res)

			fmt.Print(b.String())
		}(f)
	}

	wg.Wait()
}
