package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	fs := New(100)

	assert.Equal(t, 100, cap(fs.data))
	assert.Equal(t, 0, fs.cwdIdx)
	assert.Equal(t, "/", fs.data[0].name)
	assert.True(t, fs.data[0].isDir)
}

func TestAddDir(t *testing.T) {
	fs := New(100)
	name := "toast"
	idx := fs.AddDir(name)

	assert.Equal(t, name, fs.data[idx].name, "data should have same name as provided")
	assert.Equal(t, 0, fs.data[idx].size, "directory should start with size 0")
	assert.True(t, fs.data[idx].isDir, "directory should indicate that it is one")

}

func TestAddDirRelationship(t *testing.T) {
	fs := New(100)
	name := "toast"
	idx := fs.AddDir(name)

	cdIdx := fs.cwdIdx
	parentIdx := fs.data[idx].parent

	assert.Equal(t, 1, len(fs.data[parentIdx].children), "the number of children for the cwd should now be 1")
	assert.Equal(t, cdIdx, parentIdx, "the parent should be the cwd")
	assert.Contains(t, fs.data[parentIdx].children, idx, "the parent should contain the new data index in its slice of children")
}

func TestAddFile(t *testing.T) {
	fs := New(100)
	name := "toast"
	size := 55
	idx := fs.AddFile(name, size)

	assert.Equal(t, name, fs.data[idx].name, "data should have same name as provided")
	assert.Equal(t, 55, fs.data[idx].size, "file should have the size given")
	assert.False(t, fs.data[idx].isDir, "file should not indicate that it is a directory")
}

func TestAddFileRelationship(t *testing.T) {
	fs := New(100)
	name := "toast"
	size := 55
	idx := fs.AddFile(name, size)

	cdIdx := fs.cwdIdx
	parentIdx := fs.data[idx].parent

	assert.Equal(t, 1, len(fs.data[parentIdx].children), "the number of children for the cwd should now be 1")
	assert.Equal(t, cdIdx, parentIdx, "the parent should be the cwd")
	assert.Contains(t, fs.data[parentIdx].children, idx, "the parent should contain the new data index in its slice of children")
}