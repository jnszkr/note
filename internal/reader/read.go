package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	InvalidDateError = errors.New("invalid date")
	InvalidLineError = errors.New("invalid line")
)

type noteReader struct {
	f *os.File
	s *bufio.Scanner

	NoteReader
}

func New(path string) (NoteReader, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &noteReader{
		f: f,
		s: bufio.NewScanner(f),
	}, nil
}

func (r *noteReader) Next() bool {
	return r.s.Scan()
}

func (r *noteReader) ReadNote() (*time.Time, string, error) {
	t := r.s.Text()
	if len(t) < 26 {
		return nil, "", InvalidLineError
	}
	date, err := time.Parse(time.RFC3339, t[:25])
	if err != nil {
		return nil, "", fmt.Errorf("%v: %w", err, InvalidDateError)
	}
	msg := t[26:]
	return &date, msg, nil
}

func (r *noteReader) Close() error {
	return r.f.Close()
}
