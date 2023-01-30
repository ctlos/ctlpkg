package cmd

import (
	"gitea.dancheg97.ru/dancheg97/go-pacman/fileserver"
	"gitea.dancheg97.ru/dancheg97/go-pacman/handlers"
	"gitea.dancheg97.ru/dancheg97/go-pacman/packages"
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

	go fileserver.RunFileServer(pkgPath, viper.GetInt(`file-port`))

	packager, err := packages.Get(viper.GetString(`user`), pkgPath, viper.GetString(`repo`))
	checkErr(err)

	start := viper.GetString(`init-pkgs`)
	if start != `` {
		packager.Add(start)
	}

	err = handlers.Run(&handlers.Params{
		Port:     viper.GetInt(`grpc-port`),
		Packager: packager,
	})
	checkErr(err)
}
