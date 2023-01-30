package packages

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const pkgExt = `.pkg.tar.zst`

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

	cmd := exec.Command("bash", "-c", "yay --noconfirm --noremovemake -Sy "+packages)
	cmd.Stdout = logrus.StandardLogger().Writer()
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to execute yay for '"+packages+"': %w ", err)
	}

	cmd = exec.Command("bash", "-c", "sudo chmod a+rwx -R /home/makepkg/.cache/yay")
	cmd.Stdout = logrus.StandardLogger().Writer()
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to set proper permissions for yay build dir")
	}

	cmd = exec.Command("bash", "-c", "sudo chmod a+rwx -R /var/cache/pacman/pkg")
	cmd.Stdout = logrus.StandardLogger().Writer()
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to set proper permissions for yay build dir")
	}

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

	cmd = exec.Command("bash", "-c", "repo-add -n -q "+repo+" "+pkgs)
	cmd.Stdout = logrus.StandardLogger().Writer()
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("unable to add packages to repository")
	}

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

			err = os.WriteFile(p.PacmanCacheDir+pkgFiles.Name(), input, 0o600)
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

func (p *Packager) Search(pattern string) ([]string, error) {
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

func (p *Packager) Update() error {
	out, err := exec.Command("bash", "-c", "yay --noconfirm -Syu").Output()
	if err != nil {
		logrus.Error("yay script failed: ", string(out))
		return fmt.Errorf("unable to execute yay: %w", err)
	}
	logrus.Info("yay update script succeed: ", string(out))
	return nil
}
