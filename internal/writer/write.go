package writer

import (
	"bytes"
	"os"
	"path/filepath"
	"time"
)

type NoteWriter struct {
	f   *os.File
	now func() time.Time
}

func newWith(path string, now func() time.Time) (*NoteWriter, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &NoteWriter{
		f:   f,
		now: now,
	}, nil
}

func New(path string) (*NoteWriter, error) {
	return newWith(path, time.Now)
}

func (w *NoteWriter) WriteNote(note string) error {
	b := bytes.Buffer{}
	b.WriteString(w.now().Format(time.RFC3339))
	b.WriteString(" ")
	b.WriteString(note)
	b.WriteString("\n")

	_, err := w.f.Write(b.Bytes())
	return err
}

func (w *NoteWriter) Close() error {
	return w.f.Close()
}
