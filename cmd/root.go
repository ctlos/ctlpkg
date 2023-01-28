package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-pacman",
	Short: "Tool for managing personal pacman repository written in go.",
	Long: `
GO-PACMAN  Copyright (C) 2023  Dancheg97
This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
This is free software, and you are welcome to redistribute it
under certain conditions; watch license for details.
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
