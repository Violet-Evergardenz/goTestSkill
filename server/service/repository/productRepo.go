package repository

import (
	"myapp/model/entity"

	"gorm.io/gorm"
)

func GetAllProduct(h *gorm.DB, prods *[]entity.Products) *[]entity.Products {
	if err := h.Raw("select * from products p join brands b on p.brandId = b.brandId ").Scan(&prods).Error; err != nil {
		return nil
	}
	return prods
}

func GetProduct(h *gorm.DB, prod *entity.Products, id string) *entity.Products {
	err := h.Raw("select * from products p join brands b on p.brandId = b.brandId where productId = ?", id).Scan(&prod).Error
	if err != nil {
		return nil
	}
	return prod
}
