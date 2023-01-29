package pkg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	pkgr := Packager{
		YayCacheDir: "/home/makepkg/.cache/yay/",
	}
	err := pkgr.Add(`yay`)
	assert.NoError(t, err)
	files, err := os.ReadDir(".")
	assert.NoError(t, err)
	var found bool
	for _, file := range files {
		if file.Name() == ".pkg.zst" {
			found = true
		}
	}
	assert.True(t, found)
}
