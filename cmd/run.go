package cmd

import (
	"os"

	"gitea.dancheg97.ru/dancheg97/go-pacman/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	pgConnString    = `pg-str`
	grpcPort        = `grpc-port`
	defaultGrpcPort = 9080
	filePort        = `file-port`
	defaultFilePort = 8080
	defaultPg       = `packages`
)

func init() {
	rootCmd.Flags().String(pgConnString, defaultPg, "directory that will be used to store packages")
	viper.BindPFlag(pgConnString, rootCmd.Flags().Lookup(pgConnString))
	viper.BindEnv(pgConnString, `DIRECTORY`)

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

	err := api.Run(&api.Params{
		Port: viper.GetInt(grpcPort),
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
