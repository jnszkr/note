package searcher

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearcher(t *testing.T) {

	t.Run("Searching for 'this'", func(t *testing.T) {
		searchDir, err := filepath.Abs("./testdata")
		assert.NoError(t, err)
		var buf bytes.Buffer

		s := New(searchDir, &buf)
		s.Search("this")

		assert.Equal(t, expectedThisOut, buf.String())
	})

	t.Run("Searching for 'no match'", func(t *testing.T) {
		searchDir, err := filepath.Abs("./testdata")
		assert.NoError(t, err)
		var buf bytes.Buffer

		s := New(searchDir, &buf)
		s.Search("no match")

		assert.Equal(t, ``, buf.String())
	})

	t.Run("Searching for movie related topics", func(t *testing.T) {
		searchDir, err := filepath.Abs("./testdata")
		assert.NoError(t, err)
		var buf bytes.Buffer

		s := New(searchDir, &buf)
		s.Search("Mad Max")

		assert.Equal(t, expectedMadMaxOut, buf.String())
	})
}

const expectedThisOut = ` • 
   2021-12-24 14:03:37 [31mThis[0m is the testdata folder. Here, all the generic notes should be here.
 • movies
   2021-12-24 14:02:03 Notes in [31mthis[0m folder should contain only movie related things. 
 • movies • action
   2021-12-24 22:35:08 I heard that [31mthis[0m Mad Max movie is good 
 • music
   2021-12-24 13:55:21 [31mThis[0m note should be added in the testdata/music folder. 
              13:57:02 In [31mthis[0m folder, only music related notes are allowed 🪕. 
`

const expectedMadMaxOut = ` • movies • action
   2021-12-24 22:35:08 I heard that this [31mMad Max[0m movie is good 
`
