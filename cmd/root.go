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
		Name:        "repo",
		ShortName:   "r",
		Env:         "REPO",
		Value:       "localhost",
		Description: "📄 repository name on the web page",
	},
	{
		Cmd:         rootCmd,
		Name:        "user",
		ShortName:   "u",
		Env:         "USER",
		Value:       "makepkg",
		Description: "😀 user name in system",
	},
	{
		Cmd:         rootCmd,
		Name:        "grpc-port",
		ShortName:   "g",
		Env:         "GRPC_PORT",
		Value:       "9080",
		Type:        "int",
		Description: "🌐 gRPC API port for repository packages",
	},
	{
		Cmd:         rootCmd,
		Name:        "file-port",
		ShortName:   "f",
		Env:         "FILE_PORT",
		Value:       "8080",
		Type:        "int",
		Description: "🌐 port for static file server to access packages",
	},
	{
		Cmd:         rootCmd,
		Name:        "init-pkgs",
		ShortName:   "i",
		Env:         "INIT_PKGS",
		Description: "📦 initial packages for download",
	},
	{
		Cmd:         rootCmd,
		Name:        "logs-fmt",
		ShortName:   "l",
		Env:         "LOGS_FMT",
		Value:       "json",
		Description: "📒 output format for logs",
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
