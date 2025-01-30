package migrations

import (
	"log"
	"ultradex/database"
	"ultradex/model"
)

func init() {
	if err := createUltramanTable(); err != nil {
		log.Fatalf("Error creating Ultraman table: %v", err)
	}
}

func createUltramanTable() error {
	db, err := database.NewPostgres()
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&model.Ultraman{}); err != nil {
		return err
	}

	log.Println("ultraman table migrated successfully")

	return nil
}
