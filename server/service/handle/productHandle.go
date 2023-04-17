package handle

import (
	"myapp/model/shirt/entity"

	"myapp/server/service/repository"

	"gorm.io/gorm"
)

func GetAllProductHandle(h *gorm.DB, prods *[]entity.Products) (*[]entity.Products, error) {
	ps, err := repository.GetAllProduct(h, prods)
	return ps, err
}

func GetProductHandle(h *gorm.DB, prod *entity.Products, id string) (*entity.Products, error) {
	p, err := repository.GetProduct(h, prod, id)
	return p, err
}

func EditProductHandle(h *gorm.DB, prodNew *entity.Products) *entity.Products {
	// check any row before edit
	// return nil n msgerr
	return prodNew
}
