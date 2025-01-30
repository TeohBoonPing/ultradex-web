package imports

import (
	"fmt"

	"ultradex/model"
	"ultradex/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ImportUltraman(db *gorm.DB) error {
	records, columnIndex, err := utils.ReadCsvFile("database/csv/ultraman.csv")
	if err != nil {
		return fmt.Errorf("failed to read csv file: %w", err)
	}

	if len(records) == 0 {
		fmt.Println("no records found in the csv file")
	}

	for _, record := range records {
		height, heightErr := utils.StringToFloat64Ptr(record[columnIndex["height"]])
		if heightErr != nil {
			return fmt.Errorf("invalid height value: %w", heightErr)
		}

		weight, weightErr := utils.StringToFloat64Ptr(record[columnIndex["weight"]])
		if weightErr != nil {
			return fmt.Errorf("invalid weight value: %w", weightErr)
		}

		age, ageErr := utils.StringToUintPtr(record[columnIndex["age"]])
		if ageErr != nil {
			fmt.Println(record[columnIndex["age"]])
			return fmt.Errorf("invalid age value: %d", age)
		}

		entry := model.Ultraman{
			Name:                record[columnIndex["name"]],
			Slug:                record[columnIndex["slug"]],
			Era:                 utils.StringToPtr(record[columnIndex["era"]]),
			HumanHost:           utils.StringToPtr(record[columnIndex["human_host"]]),
			Height:              height,
			Weight:              weight,
			Age:                 age,
			HomeWorld:           utils.StringToPtr(record[columnIndex["home_world"]]),
			Race:                utils.StringToPtr(record[columnIndex["race"]]),
			FirstAppearanceYear: record[columnIndex["first_appearance_year"]],
			Description:         utils.StringToPtr(record[columnIndex["description"]]),
		}

		if err := db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "slug"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "era", "human_host", "height", "weight", "age", "home_world", "race", "first_appearance_year", "description",
			}),
		}).Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to upsert ultraman record: %w", err)
		}
	}

	return nil
}
