package controller

import (
	"ultradex/model"
	"ultradex/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UltramanController interface {
	GetUltraman(c *gin.Context) (*model.Ultraman, error)
	GetUltramanList(c *gin.Context) ([]model.Ultraman, error)
}

type ultramanController struct {
	service service.UltramanService
}

func NewUltramanController(db *gorm.DB) UltramanController {
	return &ultramanController{service: service.NewUltramanService(db)}
}

func (u *ultramanController) GetUltraman(c *gin.Context) (*model.Ultraman, error) {
	ultraman, err := u.service.FindBySlug(c.Param("slug"))
	if err != nil {
		return nil, err
	}

	return ultraman, nil
}

func (u *ultramanController) GetUltramanList(c *gin.Context) ([]model.Ultraman, error) {
	ultramans, err := u.service.FindAllUltramans()
	if err != nil {
		return nil, err
	}

	return *ultramans, nil
}
