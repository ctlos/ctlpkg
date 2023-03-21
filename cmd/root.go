package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-pacman",
	Short: "📦 Tool for managing personal pacman repository written in go.",
	Long: `📦 Tool for managing personal pacman repository written in go.

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
		Env:         "CTLPKG_REPO",
		Value:       "localhost",
		Description: "📄 repository name on the web page",
	},
	{
		Cmd:         rootCmd,
		Name:        "user",
		ShortName:   "u",
		Env:         "CTLPKG_USER",
		Value:       "makepkg",
		Description: "😀 user name in system, setted in dockerfile aswell",
	},
	{
		Cmd:         rootCmd,
		Name:        "port",
		ShortName:   "g",
		Env:         "CTLPKG_PORT",
		Value:       "8080",
		Type:        "int",
		Description: "🌐 publically exposed port for both HTTP and gRPC calls",
	},
	{
		Cmd:         rootCmd,
		Name:        "init-pkgs",
		ShortName:   "i",
		Env:         "CTLPKG_INIT_PKGS",
		Description: "📦 initial packages for download",
	},
	{
		Cmd:         rootCmd,
		Name:        "logs-fmt",
		ShortName:   "l",
		Env:         "CTLPKG_LOGS_FMT",
		Value:       "json",
		Description: "📒 output format for logs",
	},
	{
		Cmd:         rootCmd,
		Name:        "web-dir",
		ShortName:   "w",
		Env:         "CTLPKG_WEB_DIR",
		Value:       "/web",
		Description: "📂 directory with flutter web app",
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
