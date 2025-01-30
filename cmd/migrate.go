package cmd

import (
	"ultradex/database/migrations"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.Migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
