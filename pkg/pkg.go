package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Packager struct {
	PkgPath  string
	YayCache string
}

func Get(pkgPath string, yayPath string) (*Packager, error) {
	packager := &Packager{
		PkgPath:  pkgPath,
		YayCache: yayPath,
	}
	err := os.Mkdir(pkgPath+"/x86_64", os.ModePerm)
	if !os.IsExist(err) && err != nil {
		return nil, fmt.Errorf("unable to create directory for packages: %w", err)
	}
	return packager, nil
}

func (p *Packager) Add(name string) error {
	_, err := exec.Command("yay", "-Sy", name).Output()
	if err != nil {
		return fmt.Errorf("unable to execute yay for '"+name+"': %w ", err)
	}
	dir, err := os.ReadDir(p.YayCache + "/" + name)
	if err != nil {
		return fmt.Errorf("unable to find cache dir, %w", err)
	}
	var transfered bool
	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".pkg.tar.zst") {
			input, err := os.ReadFile(p.YayCache + "/" + name + "/" + file.Name())
			if err != nil {
				return fmt.Errorf("unable to read file, %w", err)
			}

			err = os.WriteFile(p.PkgPath+"/x86_64/"+file.Name(), input, 0644)
			if err != nil {
				return fmt.Errorf("unable to read file, %w", err)
			}
			transfered = true
		}
	}
	if !transfered {
		return fmt.Errorf("file was not found in build dir")
	}
	return nil
}
