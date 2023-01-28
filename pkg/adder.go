package pkg

import (
	"fmt"
	"os/exec"
)

func Add(name string) error {
	out, err := exec.Command("yay", "-Sy", name).Output()
	if err != nil {
		return fmt.Errorf("unable to get file, %w", err)
	}
	
	return nil
}
