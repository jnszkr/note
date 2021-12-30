package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jnszkr/note/internal/reader"
)

func TestFormat(t *testing.T) {
	r, err := reader.New("./testdata/.notes")
	assert.NoError(t, err)
	assert.NotNil(t, r)

	expected := `2021-12-24 13:55:21 This note should be added in the testdata/music folder. 
           13:57:02 In this folder, only music related notes are allowed 🪕. 
           13:58:00 Tycho - Easy makes me relax today! 
2021-12-29 17:14:52 Too much Xmas songs... 
           17:16:20 It is time to switch to something else 
`

	assert.Equal(t, expected, Format(r))
}
