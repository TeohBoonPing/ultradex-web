package cmd

import (
	"ultradex/database"
	"ultradex/router"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewPostgres()
		if err != nil {

		}
		router.NewRouter(db)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
