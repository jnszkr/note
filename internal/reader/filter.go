package reader

import (
	"io"
	"time"
)

type FilterFunc func(t *time.Time, note string) bool

type filterNoteReader struct {
	r       NoteReader
	filter  FilterFunc
	err     error
	bufDate *time.Time
	bufNote string

	NoteReader
}

func NewWith(path string, filter FilterFunc) (NoteReader, error) {
	r, err := New(path)
	if err != nil {
		return nil, err
	}
	return &filterNoteReader{
		r:       r,
		filter:  filter,
		err:     nil,
		bufDate: nil,
		bufNote: "",
	}, nil
}

func (f *filterNoteReader) Next() bool {
	for f.r.Next() {
		t, n, err := f.r.ReadNote()
		if err != nil {
			f.err = err
			f.bufNote = n
			f.bufDate = t
			return false
		}
		if f.filter(t, n) {
			f.bufNote = n
			f.bufDate = t
			return true
		}
	}
	f.err = io.EOF
	return false
}

func (f *filterNoteReader) ReadNote() (*time.Time, string, error) {
	if f.err == io.EOF {
		return nil, "", nil
	}
	return f.bufDate, f.bufNote, f.err
}

func (f *filterNoteReader) Close() error {
	return f.r.Close()
}
