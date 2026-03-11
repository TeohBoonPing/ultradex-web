package imports

import (
	"fmt"

	"ultradex/model"
	"ultradex/utils"

	"gorm.io/gorm"
)

func ImportUltraman(db *gorm.DB) error {
	if err := db.Exec("TRUNCATE TABLE ultramen RESTART IDENTITY").Error; err != nil {
		return fmt.Errorf("failed to truncate table: %w", err)
	}

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
			return fmt.Errorf("invalid age value: %w", ageErr)
		}

		// Parse the new stat fields
		strength, strErr := utils.StringToUintPtr(record[columnIndex["strength"]])
		if strErr != nil {
			return fmt.Errorf("invalid strength value: %w", strErr)
		}

		agility, agiErr := utils.StringToUintPtr(record[columnIndex["agility"]])
		if agiErr != nil {
			return fmt.Errorf("invalid agility value: %w", agiErr)
		}

		energy, eneErr := utils.StringToUintPtr(record[columnIndex["energy"]])
		if eneErr != nil {
			return fmt.Errorf("invalid energy value: %w", eneErr)
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
			Finisher:            utils.StringToPtr(record[columnIndex["finisher"]]),
			Color:               utils.StringToPtr(record[columnIndex["color"]]),
			Strength:            strength,
			Agility:             agility,
			Energy:              energy,
		}

		if err := db.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to insert ultraman records: %w", err)
		}
	}

	return nil
}
