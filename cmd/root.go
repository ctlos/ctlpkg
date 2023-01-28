package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-pacman",
	Short: "📦 Tool for managing personal pacman repository written in go.",
	Long: `
	📦 Tool for managing personal pacman repository written in go.

	go-pacman  Copyright (C) 2023  Dancheg97
	
	This program comes with ABSOLUTELY NO WARRANTY; for details 'use -h flag'.
	This is free software, and you are welcome to redistribute it
	under certain conditions; watch license for details.
`,
}

var flags = []Flag{
	{
		Cmd:         rootCmd,
		Name:        "yay-path",
		Env:         "YAY_PATH",
		Value:       "/home/makepkg/.cache/yay",
		Description: "📂 directory with yay cache",
	},
	{
		Cmd:         rootCmd,
		Name:        "pkg-path",
		Env:         "PKG_PATH",
		Value:       "/home/makepkg/packages",
		Description: "📂 directory to store packages",
	},
	{
		Cmd:         rootCmd,
		Name:        "grpc-port",
		Env:         "GRPC_PORT",
		Value:       "9080",
		Description: "🌐 gRPC API port for repository packages",
	},
	{
		Cmd:         rootCmd,
		Name:        "file-port",
		Env:         "FILE_PORT",
		Value:       "8080",
		Description: "🌐 port for static file server to access packages",
	},
}

func Execute() {
	for _, flag := range flags {
		AddFlag(flag)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
