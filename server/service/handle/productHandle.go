package handle

import (
	"myapp/model/shirt/entity"
	"myapp/model/shirt/request"
	"strconv"

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

func EditProductHandle(h *gorm.DB, prodNew *request.Products) (*request.Products, error) {
	// check any row before edit
	if err := repository.EditProduct(h, prodNew); err != nil {
		return nil, err
	}
	if err := repository.DelProductColor(h, strconv.Itoa(prodNew.ProductId)); err != nil {
		return nil, err
	}
	for i := 0; i < len(prodNew.ProductColors); i++ {
		if err := repository.AddProductColor(h, prodNew.ProductColors[i]); err != nil {
			return nil, err
		}
	}
	return prodNew, nil
}

func AddProductHandle(h *gorm.DB, prodNew *request.Products) error {
	if err := repository.AddProduct(h, prodNew); err != nil {
		return err
	}
	for i := 0; i < len(prodNew.ProductColors); i++ {
		if err := repository.AddProductColor(h, prodNew.ProductColors[i]); err != nil {
			return err
		}
	}
	return nil
}

func DelProductHandle(h *gorm.DB, id string) error {
	if err := repository.DelProductColor(h, id); err != nil {
		return err
	}
	if err := repository.DelProduct(h, id); err != nil {
		return err
	}
	return nil
}
