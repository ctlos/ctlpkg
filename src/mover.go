package src

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Mover interface {
	Move(src string, dst string) error
}

type ZstMover struct{}

func (z *ZstMover) Move(src string, dst string) error {
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
