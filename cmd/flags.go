package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
		Description: "😀 user name in underlying linux system",
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
		Name:        "logs-format",
		ShortName:   "f",
		Env:         "CTLPKG_LOGS_FORMAT",
		Value:       "json",
		Description: "📒 output format for logs (json/text/pretty)",
	},
	{
		Cmd:         rootCmd,
		Name:        "web-dir",
		ShortName:   "w",
		Env:         "CTLPKG_WEB_DIR",
		Value:       "/web",
		Description: "📂 directory with flutter web app",
	},
	{
		Cmd:         rootCmd,
		Name:        "api-adress",
		ShortName:   "a",
		Env:         "CTLPKG_API_ADRESS",
		Value:       "http://localhost:8080/",
		Description: "📫 adress for backend api calls via grpc-web",
	},
	{
		Cmd:         rootCmd,
		Name:        "logins",
		ShortName:   "l",
		Env:         "CTLPKG_LOGINS",
		Value:       "user|password",
		Description: "🔐 list of logins and passwords separated by '|' symbol",
	},
}

// Short description of contents for command.
type Flag struct {
	// Cobra command that we will bound our cmd to
	Cmd *cobra.Command
	// Name of command in CLI
	Name string
	// Optional short name for command, leave empty to skip short name
	ShortName string
	// Environment variable to read from
	Env string
	// Regular name for the flag
	Value string
	// Wether this value should be provided by user
	IsRequired bool
	// Leave empty if type is string: ["", "strarr", "bool"]
	Type string
	// Description for flag
	Description string
}

// Function to add new command to CLI tool.
func AddFlag(cmd Flag) {
	if cmd.Type == "" {
		cmd.Cmd.PersistentFlags().StringP(cmd.Name, cmd.ShortName, cmd.Value, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "strarr" {
		cmd.Cmd.PersistentFlags().StringArrayP(cmd.Name, cmd.ShortName, nil, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "bool" {
		cmd.Cmd.PersistentFlags().BoolP(cmd.Name, cmd.ShortName, false, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Type == "int" {
		if cmd.Value != "" {
			i, err := strconv.Atoi(cmd.Value)
			if err != nil {
				err = fmt.Errorf("value for flag "+cmd.Name+" should be int: %w", err)
				checkErr(err)
			}
			cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, i, cmd.Description)
			err = viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
			checkErr(err)
			return
		}
		cmd.Cmd.PersistentFlags().IntP(cmd.Name, cmd.ShortName, 0, cmd.Description)
		err := viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
		checkErr(err)
	}

	if cmd.Env != `` {
		err := viper.BindEnv(cmd.Name, cmd.Env)
		checkErr(err)
	}

	if cmd.IsRequired {
		err := cmd.Cmd.MarkFlagRequired(cmd.Name)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func setLogFormat() {
	switch viper.GetString(`logs-format`) {
	case `json`:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case `text`:
		logrus.SetFormatter(&logrus.TextFormatter{})
	case `pretty`:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			DisableQuote:  true,
			FullTimestamp: true,
		})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
