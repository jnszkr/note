package notes

import (
	"bufio"
	"os"
	"path/filepath"
	"time"
)

type NoteReader interface {
	Next() bool
	ReadNote() *Note
}

type noteReader struct {
	f    *os.File
	s    *bufio.Scanner
	curr *Note

	NoteReader
}

func Reader(path string) NoteReader {
	path, err := filepath.Abs(path)
	if err != nil {
		return NeverReader()
	}

	f, err := os.Open(path)
	if err != nil {
		return NeverReader()
	}
	return &noteReader{
		f: f,
		s: bufio.NewScanner(f),
	}
}

func (r *noteReader) Next() bool {
	for {
		next := r.s.Scan()
		if next {
			t := r.s.Text()
			if len(t) < 26 {
				continue // to next when row is invalid
			}
			date, err := time.Parse(time.RFC3339, t[:25])
			if err != nil {
				continue // to next when date cannot be parsed
			}
			r.curr = &Note{
				Created: date,
				Text:    t[26:],
			}
		}
		return next
	}
}

func (r *noteReader) ReadNote() *Note {
	return r.curr
}

func (r *noteReader) Close() error {
	return r.f.Close()
}

// NeverReader reader
type neverReader struct{}

func (r *neverReader) Next() bool {
	return false
}

func (r *neverReader) ReadNote() *Note {
	return &Note{}
}

func (r *neverReader) Close() error {
	return nil
}

func NeverReader() NoteReader {
	return &neverReader{}
}
