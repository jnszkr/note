package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	InvalidLineError = errors.New("invalid line")
)

type NoteReader struct {
	f    *os.File
	s    *bufio.Scanner
	once *sync.Once
}

func New(path string) (*NoteReader, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &NoteReader{
		f:    f,
		s:    bufio.NewScanner(f),
		once: &sync.Once{},
	}, nil
}

func (r *NoteReader) Next() bool {
	more := r.s.Scan()
	if !more {
		r.once.Do(func() {
			r.f.Close()
		})
	}
	return more
}

func (r *NoteReader) Read() (*time.Time, string, error) {
	t := r.s.Text()
	if len(t) < 26 {
		return nil, "", InvalidLineError
	}
	date, err := time.Parse(time.RFC3339, t[:25])
	if err != nil {
		return nil, "", fmt.Errorf("%v: %w", err, InvalidLineError)
	}
	msg := t[26:]
	return &date, msg, nil
}
