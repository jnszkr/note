package notes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	expected := `2022-02-04 23:16:07 Test data row 1
           23:16:15 Test data row 2
           23:16:22 Test data row 3
           23:16:25 Test data row 4
           23:16:26 Test data row 5
2022-02-05 23:16:28 Test data row 6
           23:16:29 Test data row 7
2022-02-06 20:16:32 Test data row 8
           20:16:35 Test data row 9
`

	assert.Equal(t, expected, Formatter()(Reader("./testdata/.notes")))
}

func TestHighlight(t *testing.T) {
	t.Run("should highlight substring", func(t *testing.T) {
		expected := "\u001B[31mThis is an example\x1b[0m text"
		assert.Equal(t, expected, Highlight("This is an example", RedColor, "This is an example text"))
	})

	t.Run("should highlight all", func(t *testing.T) {
		expected := "\u001B[31mThis is an example text\u001B[0m"
		assert.Equal(t, expected, Highlight("", RedColor, "This is an example text"))
	})
}
