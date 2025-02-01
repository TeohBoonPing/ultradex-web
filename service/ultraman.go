package service

import (
	"ultradex/model"

	"gorm.io/gorm"
)

type UltramanService interface {
	FindBySlug(slug string) (*model.Ultraman, error)
	FindAllUltramans() (*[]model.Ultraman, error)
}

type ultramanService struct {
	db *gorm.DB
}

func NewUltramanService(db *gorm.DB) UltramanService {
	return &ultramanService{db: db}
}

// FindBySlug returns the ultraman by slug
func (u *ultramanService) FindBySlug(slug string) (*model.Ultraman, error) {
	var ultraman model.Ultraman
	result := u.db.Where("slug = ?", slug).First(&ultraman)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ultraman, nil
}

// FindAllUltramans returns the list of all ultramans
func (u *ultramanService) FindAllUltramans() (*[]model.Ultraman, error) {
	var ultramans []model.Ultraman
	result := u.db.Order("id DESC").Find(&ultramans)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ultramans, nil
}
