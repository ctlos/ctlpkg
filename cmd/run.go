package cmd

import (
	"gitea.dancheg97.ru/dancheg97/go-pacman/api"
	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"gitea.dancheg97.ru/dancheg97/go-pacman/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const pkgPath = `/var/cache/pacman/pkg/`

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

	go server.RunFileServer(pkgPath, viper.GetInt(`file-port`))

	packager, err := pkg.Get(viper.GetString(`user`), pkgPath)
	checkErr(err)

	err = api.Run(&api.Params{
		Port:     viper.GetInt(`grpc-port`),
		Packager: packager,
	})
	checkErr(err)
}
