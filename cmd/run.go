package cmd

import (
	"os"

	"gitea.dancheg97.ru/dancheg97/go-pacman/api"
	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	yayPath    = `yay-path`
	defaultYay = `/home/makepkg/.cache/yay`

	pkgPath    = `pkg-path`
	defaultPkg = `/home/makepkg/packages`

	grpcPort        = `grpc-port`
	defaultGrpcPort = 9080

	filePort        = `file-port`
	defaultFilePort = 8080
)

func init() {
	rootCmd.Flags().String(pkgPath, defaultYay, "directory with yay cache")
	viper.BindPFlag(pkgPath, rootCmd.Flags().Lookup(pkgPath))
	viper.BindEnv(pkgPath, `YAY_PATH`)

	rootCmd.Flags().String(yayPath, defaultPkg, "directory to store packages")
	viper.BindPFlag(yayPath, rootCmd.Flags().Lookup(yayPath))
	viper.BindEnv(yayPath, `PKG_PATH`)

	rootCmd.Flags().Int(grpcPort, defaultGrpcPort, "gRPC API port for repository packages")
	viper.BindPFlag(grpcPort, rootCmd.Flags().Lookup(grpcPort))
	viper.BindEnv(grpcPort, `GRPC_PORT`)

	rootCmd.Flags().Int(filePort, defaultFilePort, "port for static file server to access packages")
	viper.BindPFlag(filePort, rootCmd.Flags().Lookup(filePort))
	viper.BindEnv(filePort, `FILE_PORT`)

	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run instance of go-pacman",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {
	log := logrus.StandardLogger()
	log.SetFormatter(&logrus.JSONFormatter{})

	packager, err := pkg.Get(viper.GetString(pkgPath), viper.GetString(yayPath))
	check(err, "packager")

	err = api.Run(&api.Params{
		Port:     viper.GetInt(grpcPort),
		Packager: packager,
	})
	check(err, `services`)
}

func check(err error, module string) {
	if err != nil {
		logrus.Error(`unable to load module: `, module, err)
		os.Exit(1)
	}
	logrus.Info(`module loaded successfully: `, module)
}
