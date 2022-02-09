package notes

import (
	"io"
)

type FilterFunc func(*Note) bool

type FilterResult struct {
	err    error
	r      NoteReader
	filter FilterFunc
	curr   *Note
}

func Filter(r NoteReader, filter FilterFunc) NoteReader {
	return &FilterResult{
		r:      r,
		filter: filter,
		err:    nil,
	}
}

func (f *FilterResult) Err() error {
	if f.err != nil {
		return f.err
	}
	return nil
}

func (f *FilterResult) Next() bool {
	for f.r.Next() {
		n := f.r.ReadNote()
		if f.filter(n) {
			f.curr = n
			return true
		}
	}
	f.err = io.EOF
	return false
}

func (f *FilterResult) ReadNote() *Note {
	return f.r.ReadNote()
}
