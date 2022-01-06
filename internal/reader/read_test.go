package reader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReader_Next(t *testing.T) {
	r, err := New("./testdata/readtest.txt")
	assert.NoError(t, err)

	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
}

func TestReader_ReadNote(t *testing.T) {
	r, err := New("./testdata/readtest.txt")
	assert.NoError(t, err)
	defer r.Close()

	// Line 1: 2021-12-29T09:59:25+01:00 Test message 1
	assert.True(t, r.Next())
	ts, msg, err := r.ReadNote()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:25+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 1 ", msg)

	// Line 2: 2021-12-29T09:59:29+01:00 Test message 2
	assert.True(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:29+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 2 ", msg)

	// Line 3: 2021-12-29T09:59:34+01:00
	assert.True(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.Error(t, err)
	assert.ErrorIs(t, err, InvalidLineError)
	assert.Nil(t, ts)
	assert.Equal(t, "", msg)

	// Line 4: invalid date 2021-12-29T10:21:00+01:00 Test message 4
	assert.True(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.Error(t, err)
	assert.ErrorIs(t, err, InvalidDateError)
	assert.Nil(t, ts)
	assert.Equal(t, "", msg)

	// Line 5: 2021-12-29T10:23:14+01:00 Test message 4
	assert.True(t, r.Next())
	ts, msg, err = r.ReadNote()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T10:23:14+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 4 ", msg)

	// No more lines
	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
}
