package src

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClean(t *testing.T) {
	ex, err := os.Executable()
	assert.NoError(t, err)
	exPath := filepath.Dir(ex)

	err = os.MkdirAll(exPath+"/test", 0o600)
	assert.NoError(t, err)
	err = os.Chmod(exPath+"/test", 0777)
	assert.NoError(t, err)

	cleaner := DirCleaner{}
	err = cleaner.Clean(exPath)
	assert.NoError(t, err)
	es, err := os.ReadDir(exPath)
	assert.NoError(t, err)

	for _, de := range es {
		assert.True(t, !de.IsDir())
	}
}
