package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighlight(t *testing.T) {
	assert.Equal(t,
		"\x1b[31mab\x1b[0mcd",
		Highlight("abcd", "ab", Red),
	)
	assert.Equal(t,
		"\x1b[32mThis is\x1b[0m a sentence",
		Highlight("This is a sentence", "this is", Green),
	)
	assert.Equal(t,
		"\x1b[33mEvery\x1b[0m \x1b[33mevery\x1b[0m is highlighted \x1b[33mevEry\x1b[0m",
		Highlight("Every every is highlighted evEry", "every", Yellow),
	)
	assert.Equal(t,
		"Nothing is highlighted",
		Highlight("Nothing is highlighted", "nomatch", Red),
	)
}

func TestHighlight2(t *testing.T) {
	winos = true
	assert.Equal(t,
		"no highlight",
		Highlight("no highlight", "high", Red),
	)
	winos = false
}
