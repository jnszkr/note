package notes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {

	t.Run("should not read anything when no data", func(t *testing.T) {
		f := Filter(NeverReader(), func(_ *Note) bool {
			return true
		})

		assert.False(t, f.Next())
		assert.False(t, f.Next())
	})

	t.Run("should return data", func(t *testing.T) {
		f := Filter(Reader("./testdata/.notes"), func(_ *Note) bool {
			return true
		})

		expected01 := &Note{Created: mustParseTime("2022-02-04T23:16:07+01:00"), Text: "Test data row 1"}
		assert.True(t, f.Next())
		assert.Equal(t, expected01, f.ReadNote())
		assert.Equal(t, expected01, f.ReadNote())
		assert.Equal(t, expected01, f.ReadNote())
		assert.Equal(t, expected01, f.ReadNote())

		expected02 := &Note{Created: mustParseTime("2022-02-04T23:16:15+01:00"), Text: "Test data row 2"}
		assert.True(t, f.Next())
		assert.Equal(t, expected02, f.ReadNote())
		assert.Equal(t, expected02, f.ReadNote())

		expected03 := &Note{Created: mustParseTime("2022-02-04T23:16:22+01:00"), Text: "Test data row 3"}
		assert.True(t, f.Next())
		assert.Equal(t, expected03, f.ReadNote())

		assert.True(t, f.Next())  // 4
		assert.True(t, f.Next())  // 5
		assert.True(t, f.Next())  // 6
		assert.True(t, f.Next())  // 7
		assert.True(t, f.Next())  // 8
		assert.True(t, f.Next())  // 9
		assert.False(t, f.Next()) // no more lines
		assert.False(t, f.Next())
		assert.False(t, f.Next())
	})

	t.Run("should return filtered data", func(t *testing.T) {
		f := Filter(Reader("./testdata/.notes"), func(n *Note) bool {
			return n.Text[14] == '3' || n.Text[14] == '2'
		})

		expected01 := &Note{Created: mustParseTime("2022-02-04T23:16:15+01:00"), Text: "Test data row 2"}
		assert.True(t, f.Next())
		assert.Equal(t, expected01, f.ReadNote())

		expected02 := &Note{Created: mustParseTime("2022-02-04T23:16:22+01:00"), Text: "Test data row 3"}
		assert.True(t, f.Next())
		assert.Equal(t, expected02, f.ReadNote())

		assert.False(t, f.Next()) // no more lines
	})

	t.Run("should read nothing if file does not exist", func(t *testing.T) {
		f := Filter(Reader("./testdata/doesnotexist/.notes"), func(_ *Note) bool {
			return true
		})

		assert.False(t, f.Next())
		assert.False(t, f.Next())
	})
}

func mustParseTime(val string) time.Time {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		panic("must parse time")
	}
	return t
}
