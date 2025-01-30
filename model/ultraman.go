package model

import (
	"ultradex/database"

	"gorm.io/gorm"
)

type Ultraman struct {
	gorm.Model
	Name                string   `json:"name" gorm:"type:varchar(255);not null"`
	Slug                string   `json:"slug" gorm:"type:varchar(255);unique;not null"`
	Era                 *string  `json:"era" gorm:"type:varchar(100);null"`
	HumanHost           *string  `json:"human_host" gorm:"type:varchar(255);null"`
	Height              *float64 `json:"height" gorm:"type:decimal(5,2);null"`
	Weight              *float64 `json:"weight" gorm:"type:decimal(7,2);null"`
	Age                 *uint    `json:"age" gorm:"type:int;null"`
	HomeWorld           *string  `json:"home_world" gorm:"type:varchar(255);null"`
	Race                *string  `json:"race" gorm:"type:varchar(100);null"`
	FirstAppearanceYear string   `json:"first_appearance_year" gorm:"type:char(4);not null"`
	Description         *string  `json:"description" gorm:"type:text;null"`
}

func FindAll() ([]Ultraman, error) {
	db, err := database.NewPostgres()
	if err != nil {
		return nil, err
	}

	var ultraman []Ultraman
	result := db.Find(&ultraman)
	return ultraman, result.Error
}
