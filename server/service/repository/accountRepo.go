package repository

import (
	"myapp/model/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllProduct(c echo.Context, h *gorm.DB, p *[]entity.Products) (*[]entity.Products, error) {
	if err := h.Raw("select * from products p join brands b on p.brandId = b.brandId ").Scan(p).Error; err != nil {
		return p, c.NoContent(http.StatusNoContent)
	}
	return p, nil
}
