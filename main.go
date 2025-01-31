package main

import (
	"fmt"
	"ultradex/config"
	"ultradex/database"
	"ultradex/database/imports"
	"ultradex/database/migrations"
	"ultradex/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	migrations.Migrate(db)
	imports.ImportUltraman(db)

	router.NewRouter(db)
}
