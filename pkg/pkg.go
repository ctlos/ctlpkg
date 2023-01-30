package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
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

func (p *Packager) Add(packages string) (*string, error) {
	var output string

	logrus.Info("executing yay for packages: ", packages)
	out, err := exec.Command("bash", "-c", "yay --noconfirm --noremovemake -Sy "+packages).Output()
	if err != nil {
		logrus.Error("yay script failed: ", string(out))
		return nil, fmt.Errorf("unable to execute yay for '"+packages+"': %w ", err)
	}
	logrus.Info("yay script succeed: ", string(out))
	output += string(out)

	des, err := os.ReadDir(p.YayCacheDir)
	if err != nil {
		return nil, fmt.Errorf("unable to find yay cache dir, %w", err)
	}
	for _, de := range des {
		if de.IsDir() {
			err := p.processPackageDir(de.Name())
			if err != nil {
				return nil, err
			}
		}
	}

	repo := p.PacmanCacheDir + "/" + p.RepoName + ".db.tar.gz"
	pkgs := p.PacmanCacheDir + "/*.pkg.tar.zst"
	out, err = exec.Command("bash", "-c", "repo-add -n -q "+repo+" "+pkgs).Output()
	if err != nil {
		logrus.Error("repo-add script failed: ", string(out))
		return nil, fmt.Errorf("unable to add packages to repository")
	}
	logrus.Info("repo-add script succeed: ", string(out))

	output += string(out)
	return &output, nil
}

func (p *Packager) processPackageDir(pkg string) error {
	// LATER ADD GITEA HOOK TO MIRROR SOURCE CODE

	logrus.Info("processing cached yay package: ", pkg)
	ddes, err := os.ReadDir(p.YayCacheDir + pkg)
	if err != nil {
		return fmt.Errorf("unable to find pkg cache dir: %s, %w", pkg, err)
	}

	for _, pkgFiles := range ddes {
		if strings.HasSuffix(pkgFiles.Name(), ".pkg.tar.zst") {
			logrus.Info("moving package file: ", pkgFiles.Name())

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
		logrus.Error("unable to remove cache: ", pkg)
		return fmt.Errorf("unable to remove cached dir")
	}
	logrus.Info("yay cache cleaned for directory: ", pkg)
	return nil
}
