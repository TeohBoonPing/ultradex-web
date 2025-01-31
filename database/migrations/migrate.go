package migrations

import (
	"ultradex/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	_ = db.AutoMigrate(&model.Ultraman{})
}
