package searcher

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearcher(t *testing.T) {

	searchDir, err := filepath.Abs("./testdata")
	assert.NoError(t, err)

	t.Run("Should return with results", func(t *testing.T) {
		t.Run("from subfolders", func(t *testing.T) {
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("this", true)
			assert.Equal(t, expectedThisOutRecursive, buf.String())
		})

		t.Run("from only the main folder", func(t *testing.T) {
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("this", false)
			assert.Equal(t, expectedThisOut, buf.String())
		})
	})

	t.Run("Should return no result", func(t *testing.T) {
		t.Run("from subfolders", func(t *testing.T) {
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("no match", true)
			assert.Equal(t, "", buf.String())
		})
		t.Run("from only the main folder", func(t *testing.T) {
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("Mad Max", false)
			assert.Equal(t, "", buf.String())
		})
		t.Run("from a folder that has no notes", func(t *testing.T) {
			searchDir, err := filepath.Abs("./testdata/nothinghere")
			assert.NoError(t, err)
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("Mad Max", false)
			assert.Equal(t, "", buf.String())
		})
		t.Run("from a folder that does not exist", func(t *testing.T) {
			searchDir, err := filepath.Abs("./testdata/this does not exist")
			assert.NoError(t, err)
			var buf bytes.Buffer
			s := New(searchDir, &buf)
			s.Search("Mad Max", false)
			assert.Equal(t, "", buf.String())
		})
	})
}

const expectedThisOutRecursive = ` • 
   2021-12-24 14:03:37 [31mThis[0m is the testdata folder. Here, all the generic notes should be here.
 • movies
   2021-12-24 14:02:03 Notes in [31mthis[0m folder should contain only movie related things. 
 • movies • action
   2021-12-24 22:35:08 I heard that [31mthis[0m Mad Max movie is good 
 • music
   2021-12-24 13:55:21 [31mThis[0m note should be added in the testdata/music folder. 
              13:57:02 In [31mthis[0m folder, only music related notes are allowed 🪕. 
`

const expectedThisOut = ` • 
   2021-12-24 14:03:37 [31mThis[0m is the testdata folder. Here, all the generic notes should be here.
`
