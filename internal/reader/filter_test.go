package reader

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilterNoteReader_NextWithNoFilter(t *testing.T) {
	var filter = func(t *time.Time, note string) bool {
		return true
	}
	r, err := NewWith("./testdata/filtertest.txt", filter)
	assert.NoError(t, err)

	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next()) // last line
	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
}

func TestFilterNoteReader_ReadNote(t *testing.T) {
	r, err := NewWith("./testdata/filtertest.txt", func(t *time.Time, note string) bool {
		return strings.Contains(note, "3")
	})
	assert.NoError(t, err)

	assert.True(t, r.Next())
	ts, msg, err := r.ReadNote()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:29+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 3", msg)

	assert.True(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:39+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 33", msg)

	assert.False(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.NoError(t, err)
	assert.Nil(t, ts)
	assert.Equal(t, "", msg)

	assert.False(t, r.Next())
}
