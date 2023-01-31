package src

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

type Caller interface {
	Execute(cmd string) error
}

// Interface for os bash calls.
type BashCaller struct{}

// Executes bash call and loggs output to logrus.
func (c *BashCaller) Execute(cmd string) error {
	commad := exec.Command("bash", "-c", cmd)

	commad.Stdout = logrus.StandardLogger().Writer()
	commad.Stderr = logrus.StandardLogger().Writer()

	err := commad.Run()
	if err != nil {
		return err
	}
	return nil
}
