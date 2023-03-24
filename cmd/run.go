package cmd

import (
	"dancheg97.ru/dancheg97/ctlpkg/cmd/service"
	"dancheg97.ru/dancheg97/ctlpkg/cmd/utils"
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
		pkgPath   = "/var/cache/pacman/pkg"
		yayPath   = "/home/" + viper.GetString("user") + "/.cache/yay"
		port      = viper.GetInt("port")
		repoName  = viper.GetString("repo")
		webPath   = viper.GetString("web-dir")
		initPkgs  = viper.GetString(`init-pkgs`)
		apiAdress = viper.GetString(`api-adress`)
	)

	setLogFormat()

	helper := &utils.OsHelper{}

	err := helper.ReplaceFileString(
		webPath+`/main.dart.js`,
		`http://localhost:8080/`,
		apiAdress,
	)
	checkErr(err)

	err = helper.Execute("yay -Sy --noconfirm " + initPkgs)
	checkErr(err)

	err = helper.FormDb(yayPath, pkgPath, repoName)
	checkErr(err)

	err = service.Run(&service.Params{
		Port:     port,
		PkgPath:  pkgPath,
		YayPath:  yayPath,
		WebPath:  webPath,
		RepoName: repoName,
		Packager: helper,
	})
	checkErr(err)
}
