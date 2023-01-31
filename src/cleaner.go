package src

import (
	"os"
)

type Cleaner interface {
	Clean(dir string) error
}

// Interface representing directory cleaner.
type DirCleaner struct{}

// Removes all subdirectories and it's contents from directory.
func (d *DirCleaner) Clean(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			err := os.RemoveAll(dir + "/" + e.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
