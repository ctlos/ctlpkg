package cmd

import (
	"os"

	"gitea.dancheg97.ru/dancheg97/go-pacman/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	pgConnString = `pg-str`
	port         = `port`
	defaultPg    = `packages`
	defaultPort  = 9080
)

func init() {
	rootCmd.Flags().String(pgConnString, defaultPg, "directory that will be used to store packages")
	viper.BindPFlag(pgConnString, rootCmd.Flags().Lookup(pgConnString))
	viper.BindEnv(pgConnString, `DIRECTORY`)

	rootCmd.Flags().Int(port, defaultPort, "gRPC API port for repository packages")
	viper.BindPFlag(port, rootCmd.Flags().Lookup(port))
	viper.BindEnv(port, `GRPC_PORT`)

	rootCmd.Flags().Int(port, defaultPort, "port for static file server to access packages")
	viper.BindPFlag(port, rootCmd.Flags().Lookup(port))
	viper.BindEnv(port, `FILE_PORT`)

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
		Port: viper.GetInt(port),
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
