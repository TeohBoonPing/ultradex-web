package migrations

import (
	"ultradex/database"

	"github.com/go-gormigrate/gormigrate/v2"
)

var _migrations = []*gormigrate.Migration{}

func Migrate() error {
	db, err := database.NewPostgres()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, _migrations)
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
