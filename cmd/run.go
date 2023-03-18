package cmd

import (
	"dancheg97.ru/dancheg97/ctlpkg/services"
	"dancheg97.ru/dancheg97/ctlpkg/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "🚀 Run instance of go-pacman",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	var (
		pkgPath  = "/var/cache/pacman/pkg"
		yayPath  = "/home/" + viper.GetString("user") + "/.cache/yay"
		webPath  = viper.GetString("web-dir")
		httpPort = viper.GetInt("http-port")
		grpcPort = viper.GetInt("grpc-port")
		repoName = viper.GetString("repo")
		initPkgs = viper.GetString(`init-pkgs`)
	)

	setLogFormat()

	helper := &src.OsHelper{}

	err := helper.Execute("yay -Sy --noconfirm " + initPkgs)
	checkErr(err)

	err = helper.FormDb(yayPath, pkgPath, repoName)
	checkErr(err)

	services.Run(&services.Params{
		HttpPort: httpPort,
		GrpcPort: grpcPort,
		PkgPath:  pkgPath,
		YayPath:  yayPath,
		WebPath:  webPath,
		RepoName: repoName,
		Packager: &src.OsHelper{},
	})
	checkErr(err)
}
