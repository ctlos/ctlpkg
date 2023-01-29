package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Packager struct {
	RepoName       string
	YayCacheDir    string
	PacmanCacheDir string
}

func Get(user string, pkgDir string, repoName string) (*Packager, error) {
	return &Packager{
		RepoName:       repoName,
		YayCacheDir:    `/home/` + user + `/.cache/yay/`,
		PacmanCacheDir: pkgDir,
	}, nil
}

func (p *Packager) Add(name string) error {
	_, err := exec.Command("bash", "-c", "yay --noconfirm --noremovemake -Sy "+name).Output()
	if err != nil {
		return fmt.Errorf("unable to execute yay for '"+name+"': %w ", err)
	}
	des, err := os.ReadDir(p.YayCacheDir)
	if err != nil {
		return fmt.Errorf("unable to find yay cache dir, %w", err)
	}
	for _, de := range des {
		if de.IsDir() {
			err := p.processPackageDir(de.Name())
			if err != nil {
				return err
			}
		}
	}
	repo := p.PacmanCacheDir + "/" + p.RepoName + ".db.tar.gz"
	pkgs := p.PacmanCacheDir + "/*.pkg.tar.zst"
	_, err = exec.Command("bash", "-c", "repo-add -n -q "+repo+" "+pkgs).Output()
	return err
}

func (p *Packager) processPackageDir(pkg string) error {
	// LATER ADD GITEA HOOK TO MIRROR SOURCE CODE
	ddes, err := os.ReadDir(p.YayCacheDir + pkg)
	if err != nil {
		return fmt.Errorf("unable to find pkg cache dir: %s, %w", pkg, err)
	}
	for _, pkgFiles := range ddes {
		if strings.HasSuffix(pkgFiles.Name(), ".pkg.tar.zst") {
			input, err := os.ReadFile(p.YayCacheDir + pkg + "/" + pkgFiles.Name())
			if err != nil {
				return fmt.Errorf("unable to read file, %w", err)
			}

			err = os.WriteFile(p.PacmanCacheDir+pkgFiles.Name(), input, 0600)
			if err != nil {
				return fmt.Errorf("unable to read file, %w", err)
			}
		}
	}
	err = os.RemoveAll(p.YayCacheDir + pkg)
	if err != nil {
		return fmt.Errorf("unable to remove cached dir")
	}
	return nil
}
