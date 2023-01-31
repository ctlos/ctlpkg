package src

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMover(t *testing.T) {
	ex, err := os.Executable()
	assert.NoError(t, err)
	exPath := filepath.Dir(ex)

	err = os.MkdirAll(exPath+"/first", 0777)
	assert.NoError(t, err)

	err = os.MkdirAll(exPath+"/second", 0777)
	assert.NoError(t, err)

	err = os.WriteFile(exPath+"/first/nani.pkg.tar.zst", []byte("nani"), 0o600)
	assert.NoError(t, err)

	mover := ZstMover{}
	err = mover.Move(exPath+"/first", exPath+"/second")
	assert.NoError(t, err)
	_, err = os.Stat(exPath + "/second/nani.pkg.tar.zst")
	assert.NoError(t, err)
}
