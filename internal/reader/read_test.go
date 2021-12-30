package reader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReader_Next(t *testing.T) {
	r, err := New("./testdata/.notes")
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

func TestReader_Read(t *testing.T) {
	r, err := New("./testdata/.notes")
	assert.NoError(t, err)

	assert.True(t, r.Next())
	ts, msg, err := r.Read()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:25+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 1 ", msg)

	assert.True(t, r.Next())
	ts, msg, err = r.Read()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T09:59:29+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 2 ", msg)

	assert.True(t, r.Next())
	ts, msg, err = r.Read()
	assert.Error(t, err)
	assert.ErrorIs(t, err, InvalidLineError)
	assert.Nil(t, ts)
	assert.Equal(t, "", msg)

	assert.True(t, r.Next())
	ts, msg, err = r.Read()
	assert.Error(t, err)
	assert.ErrorIs(t, err, InvalidLineError)
	assert.Nil(t, ts)
	assert.Equal(t, "", msg)

	assert.True(t, r.Next())
	ts, msg, err = r.Read()
	assert.NoError(t, err)
	assert.Equal(t, "2021-12-29T10:23:14+01:00", ts.Format(time.RFC3339))
	assert.Equal(t, "Test message 4 ", msg)

	assert.False(t, r.Next())
	assert.False(t, r.Next())
	assert.False(t, r.Next())
}
