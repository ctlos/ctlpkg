package cmd

import (
	"gitea.dancheg97.ru/dancheg97/go-pacman/services"
	"gitea.dancheg97.ru/dancheg97/go-pacman/src"
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

	helper := &src.OsHelper{}

	err := helper.Execute("yay -Sy --noconfirm " + viper.GetString(`init-pkgs`))
	checkErr(err)

	services.Run(&services.Params{
		FilePort: viper.GetInt("file-port"),
		GrpcPort: viper.GetInt("grpc-port"),
		PkgPath:  pkgPath,
		YayPath:  "/home/" + viper.GetString("user") + "/.cache/yay",
		RepoName: viper.GetString("repo"),
		Packager: &src.OsHelper{},
	})
	checkErr(err)
}
