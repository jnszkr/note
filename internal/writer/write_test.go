package writer

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// chtmpdir changes the working directory to a new temporary directory and
// provides a cleanup function.
func chtmpdir(t *testing.T) func() {
	oldwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	d, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	if err := os.Chdir(d); err != nil {
		t.Fatalf("chtmpdir: %v", err)
	}
	return func() {
		if err := os.Chdir(oldwd); err != nil {
			t.Fatalf("chtmpdir: %v", err)
		}
		os.RemoveAll(d)
	}
}

func TestNoteWriter_WriteNote(t *testing.T) {
	defer chtmpdir(t)()

	now := time.Unix(1641039682, 0)
	note := "Test note"

	w, err := newWith("./.notes", func() time.Time {
		return now
	})
	assert.NoError(t, err)
	err = w.WriteNote(note)
	assert.NoError(t, err)
	err = w.WriteNote(note)
	assert.NoError(t, err)
	w.Close()

	file, err := os.Open("./.notes")
	assert.NoError(t, err)
	c, err := io.ReadAll(file)
	assert.NoError(t, err)

	expected := fmt.Sprintf("%s %s\n%s %s\n", now.Format(time.RFC3339), note, now.Format(time.RFC3339), note)

	assert.Equal(t, expected, string(c))
}
