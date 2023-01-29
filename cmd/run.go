package cmd

import (
	"gitea.dancheg97.ru/dancheg97/go-pacman/api"
	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run instance of go-pacman",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	log := logrus.StandardLogger()
	log.SetFormatter(&logrus.TextFormatter{})

	packager, err := pkg.Get(viper.GetString(`user`))
	checkErr(err)

	err = api.Run(&api.Params{
		Port:     viper.GetInt(`grpc-port`),
		Packager: packager,
	})
	checkErr(err)
}
