package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Short description of contents for command
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

// Function to add new command to CLI tool
func AddFlag(cmd Flag) {
	if cmd.Type == "" {
		cmd.Cmd.PersistentFlags().StringP(cmd.Name, cmd.ShortName, cmd.Value, cmd.Description)
		viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
	}

	if cmd.Type == "strarr" {
		cmd.Cmd.PersistentFlags().StringArrayP(cmd.Name, cmd.ShortName, nil, cmd.Description)
		viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
	}

	if cmd.Type == "bool" {
		cmd.Cmd.PersistentFlags().BoolP(cmd.Name, cmd.ShortName, false, cmd.Description)
		viper.BindPFlag(cmd.Name, cmd.Cmd.PersistentFlags().Lookup(cmd.Name))
	}

	if cmd.IsRequired {
		cmd.Cmd.MarkFlagRequired(cmd.Name)
	}
}

func checkErr(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
