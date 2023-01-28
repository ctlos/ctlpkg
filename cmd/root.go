package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "regen",
	Short: "Tool for managing personal pacman repository.",
	Long: `
REGEN  Copyright (C) 2023  Dancheg97
This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
This is free software, and you are welcome to redistribute it
under certain conditions; type 'show c' for details.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
