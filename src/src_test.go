package src

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	ex, err := os.Executable()
	assert.NoError(t, err)
	exPath := filepath.Dir(ex)

	err = os.MkdirAll(exPath+"/first", 0o777)
	assert.NoError(t, err)

	err = os.MkdirAll(exPath+"/second", 0o777)
	assert.NoError(t, err)

	err = os.WriteFile(exPath+"/first/nani.pkg.tar.zst", []byte("nani"), 0o600)
	assert.NoError(t, err)

	mover := OsHelper{}
	err = mover.Move(exPath+"/first", exPath+"/second")
	assert.NoError(t, err)
	_, err = os.Stat(exPath + "/second/nani.pkg.tar.zst")
	assert.NoError(t, err)
}

func TestClean(t *testing.T) {
	ex, err := os.Executable()
	assert.NoError(t, err)
	exPath := filepath.Dir(ex)

	err = os.MkdirAll(exPath+"/test", 0o600)
	assert.NoError(t, err)
	err = os.Chmod(exPath+"/test", 0o777)
	assert.NoError(t, err)

	cleaner := OsHelper{}
	err = cleaner.Clean(exPath)
	assert.NoError(t, err)
	es, err := os.ReadDir(exPath)
	assert.NoError(t, err)

	for _, de := range es {
		assert.True(t, !de.IsDir())
	}
}

func TestExecute(t *testing.T) {
	c := OsHelper{}
	err := c.Execute("ls")
	assert.NoError(t, err)
}

func TestSear(t *testing.T) {
	ex, err := os.Executable()
	assert.NoError(t, err)
	exPath := filepath.Dir(ex)

	err = os.MkdirAll(exPath+"/search", 0o777)
	assert.NoError(t, err)

	err = os.WriteFile(exPath+"/search/nani.pkg.tar.zst", []byte("nani"), 0o600)
	assert.NoError(t, err)

	c := OsHelper{}
	rez, err := c.Search(exPath+"/search", "nan")
	assert.NoError(t, err)
	assert.Contains(t, rez, "nani")
}
