package cmd

import (
	"fmt"
	"ultradex/database"
	"ultradex/database/imports"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewPostgres()
		if err != nil {
			fmt.Println(err)
		}

		err = imports.ImportUltraman(db)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
