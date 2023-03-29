package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type OsHelper struct{}

// Executes bash call and loggs output to logrus.
func (o *OsHelper) Execute(cmd string) error {
	logrus.Info(`executing system call: `, cmd)

	commad := exec.Command("bash", "-c", cmd)

	commad.Stdout = logrus.StandardLogger().Writer()
	commad.Stderr = logrus.StandardLogger().Writer()

	return commad.Run()
}

func (o *OsHelper) Call(cmd string) (string, error) {
	logrus.Info(`executing system call: `, cmd)

	command := exec.Command("bash", "-c", cmd)

	out := bytes.Buffer{}
	command.Stdout = &out
	command.Stderr = &out

	err := command.Run()
	if err != nil {
		return ``, err
	}
	return out.String(), nil
}

// Moves all .zst files from one dir to another.
func (o *OsHelper) Move(src string, dst string) error {
	des, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	err = o.Execute("sudo chmod a+rwx -R " + src)
	if err != nil {
		return err
	}
	err = o.Execute("sudo chmod a+rwx -R " + dst)
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
			logrus.Info("moved file, success: ", de.Name())
		}
	}

	return nil
}

// Removes all subdirectories and it's contents from directory.
func (o *OsHelper) Clean(dir string) error {
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

func (o *OsHelper) FormDb(yay string, pkg string, repo string) error {
	entires, err := os.ReadDir(yay)
	if err != nil {
		return err
	}
	for _, de := range entires {
		if de.IsDir() {
			err := o.Move(yay+"/"+de.Name(), pkg)
			if err != nil {
				return err
			}
		}
	}
	repoPath := pkg + "/" + repo + ".db.tar.gz"
	pkgsPath := pkg + "/*.pkg.tar.zst"
	err = o.Execute("sudo repo-add -n -q " + repoPath + " " + pkgsPath)
	return err
}

func (o *OsHelper) ParsePkgInfo(inp string) map[string]string {
	logrus.Info("parsing package info: ", inp)
	out := map[string]string{}
	for _, s := range strings.Split(inp, "\n") {
		if strings.HasPrefix(s, " ") {
			continue
		}
		splitted := strings.Split(s, ":")
		correctdElementAmount := 2
		if len(splitted) < correctdElementAmount {
			continue
		}
		out[strings.TrimSpace(splitted[0])] = strings.TrimSpace(splitted[1])
	}
	return out
}

func (o *OsHelper) ParsePackages(inp string) []string {
	out := []string{}
	for _, s := range strings.Split(inp, "\n") {
		if inp == "" {
			continue
		}
		splitted := strings.Split(s, " ")
		out = append(out, splitted[0])
	}
	return out
}

func (o *OsHelper) ReplaceFileString(file string, old string, new string) error {
	contents, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	replaced := strings.ReplaceAll(string(contents), old, new)
	return os.WriteFile(file, []byte(replaced), 0o600)
}
