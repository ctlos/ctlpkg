package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	MigrConnStr   = `migr-pg`
	MigrationsDir = `migr-dir`
)

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().String(MigrConnStr, "", "connection string for postgres")
	viper.BindPFlag(MigrConnStr, migrateCmd.Flags().Lookup(MigrConnStr))
	viper.BindEnv(MigrConnStr, `MIGR_CONN_STR`)

	migrateCmd.Flags().String(MigrationsDir, "migrations/", "connection string for postgres")
	viper.BindPFlag(MigrationsDir, migrateCmd.Flags().Lookup(MigrationsDir))
	viper.BindEnv(MigrationsDir, `MIGRATIONS_DIR`)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate to new schema, provide migrations and pg connection string",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := goose.OpenDBWithDriver(`pgx`, viper.GetString(MigrConnStr))
		check(err, `db_connection`)

		err = goose.Up(db, viper.GetString(MigrationsDir))
		check(err, `goose_migrator`)
	},
}
