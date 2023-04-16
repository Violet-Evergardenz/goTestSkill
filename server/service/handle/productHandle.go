package handle

import (
	"myapp/model/entity"

	"myapp/server/service/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllProductHandle(c echo.Context, h *gorm.DB) *[]entity.Products {
	var product []entity.Products
	repository.GetAllProduct(h, &product)
	return &product
}

func GetProductHandle(c echo.Context, h *gorm.DB) *entity.Products {
	id := c.Param("id")
	var product entity.Products
	repository.GetProduct(h, &product, id)
	return &product
}
