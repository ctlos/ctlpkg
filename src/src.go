package src

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type OsHelper struct{}

// Executes bash call and loggs output to logrus.
func (c *OsHelper) Execute(cmd string) error {
	commad := exec.Command("bash", "-c", cmd)

	commad.Stdout = logrus.StandardLogger().Writer()
	commad.Stderr = logrus.StandardLogger().Writer()

	err := commad.Run()
	if err != nil {
		return err
	}
	return nil
}

// Searches for all .zst files in dir and returns back list.
func (s *OsHelper) Search(dir string, pattern string) ([]string, error) {
	var packages []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("unable to scan file: %w", err)
		}
		if strings.HasSuffix(info.Name(), ".pkg.tar.zst") && strings.Contains(info.Name(), pattern) {
			packages = append(packages, strings.Replace(info.Name(), ".pkg.tar.zst", ``, 1))
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf(`unable to search: %w`, err)
	}
	return packages, nil
}

// Moves all .zst files from one dir to another.
func (z *OsHelper) Move(src string, dst string) error {
	des, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, de := range des {
		if strings.HasSuffix(de.Name(), ".pkg.tar.zst") {
			logrus.Info("moving package file: ", de.Name())

			input, err := os.ReadFile(src + "/" + de.Name())
			if err != nil {
				return err
			}

			err = os.WriteFile(dst+"/"+de.Name(), input, 0o600)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Removes all subdirectories and it's contents from directory.
func (d *OsHelper) Clean(dir string) error {
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
