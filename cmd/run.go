package cmd

import (
	"gitea.dancheg97.ru/dancheg97/go-pacman/services"
	"gitea.dancheg97.ru/dancheg97/go-pacman/src"
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
	var (
		pkgPath  = `/var/cache/pacman/pkg/`
		yayPath  = "/home/" + viper.GetString("user") + "/.cache/yay"
		filePort = viper.GetInt("file-port")
		grpcPort = viper.GetInt("grpc-port")
		repoName = viper.GetString("repo")
		initPkgs = viper.GetString(`init-pkgs`)
	)

	helper := &src.OsHelper{}

	err := helper.Execute("yay -Sy --noconfirm " + initPkgs)
	checkErr(err)

	err = helper.FormDb(yayPath, pkgPath, repoName)
	checkErr(err)

	services.Run(&services.Params{
		FilePort: filePort,
		GrpcPort: grpcPort,
		PkgPath:  pkgPath,
		YayPath:  yayPath,
		RepoName: repoName,
		Packager: &src.OsHelper{},
	})
	checkErr(err)
}
