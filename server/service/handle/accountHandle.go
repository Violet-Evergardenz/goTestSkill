package handle

import (
	"myapp/model/entity"

	"myapp/server/service/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllProductHandle(c echo.Context, h *gorm.DB) (*[]entity.Products, error) {
	var product []entity.Products
	repository.GetAllProduct(c, h, &product)
	return &product, nil
}
