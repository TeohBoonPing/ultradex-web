package controller

import (
	"ultradex/model"
	"ultradex/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UltramanController interface {
	ListUltramans(c *gin.Context) ([]model.Ultraman, error)
}

type ultramanController struct {
	service service.UltramanService
}

func NewUltramanController(db *gorm.DB) UltramanController {
	return &ultramanController{service: service.NewUltramanService(db)}
}

func (u *ultramanController) ListUltramans(c *gin.Context) ([]model.Ultraman, error) {
	ultramans, err := u.service.FindAllUltramans()
	if err != nil {
		return nil, err
	}

	return *ultramans, nil
}
