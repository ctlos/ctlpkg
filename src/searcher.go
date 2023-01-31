package src

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Searcher struct{}

func (s *Searcher) Search(pattern string) ([]string, error) {
	var packages []string
	err := filepath.Walk(p.PacmanCacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("unable to scan file: %w", err)
		}
		if strings.HasSuffix(info.Name(), pkgExt) && strings.Contains(info.Name(), pattern) {
			packages = append(packages, strings.Replace(info.Name(), pkgExt, ``, 1))
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf(`unable to search: %w`, err)
	}
	return packages, nil
}
