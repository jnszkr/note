package files

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	fs := Find("./testdata", "somefile.txt")
	filepath, next := <-fs
	fi, err := os.Stat(filepath)
	assert.NoError(t, err)
	assert.NotContains(t, "lvl1/lvl2/lvl3", filepath)
	assert.Equal(t, "somefile.txt", fi.Name())
	assert.False(t, fi.IsDir())
	assert.True(t, next)

	filepath, next = <-fs
	assert.Equal(t, "", filepath)
	assert.False(t, next)
}

func TestRecursiveFind(t *testing.T) {
	fs := RecursiveFind("./testdata", "somefile.txt")
	filepath, next := <-fs
	fi, err := os.Stat(filepath)
	assert.NoError(t, err)
	assert.Equal(t, "somefile.txt", fi.Name())
	assert.False(t, fi.IsDir())
	assert.True(t, next)

	filepath, next = <-fs
	fi, err = os.Stat(filepath)
	assert.NoError(t, err)
	assert.Equal(t, "somefile.txt", fi.Name())
	assert.False(t, fi.IsDir())
	assert.True(t, next)

	filepath, next = <-fs
	assert.Equal(t, "", filepath)
	assert.False(t, next)
}
