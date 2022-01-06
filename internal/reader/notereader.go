package reader

import (
	"io"
	"time"
)

type NoteReader interface {
	Next() bool
	ReadNote() (*time.Time, string, error)

	io.Closer
}
