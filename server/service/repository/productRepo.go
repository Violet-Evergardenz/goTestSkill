package repository

import (
	"myapp/model/shirt/entity"
	"myapp/model/shirt/request"
	"time"

	"gorm.io/gorm"
)

func GetAllProduct(h *gorm.DB, prods *[]entity.Products) (*[]entity.Products, error) {
	if err := h.Raw("select * from products p join brands b on p.brandId = b.brandId ").Scan(&prods).Error; err != nil {
		return nil, err
	}
	return prods, nil
}

func GetProduct(h *gorm.DB, prod *entity.Products, id string) (*entity.Products, error) {
	err := h.Raw("select * from products p join brands b on p.brandId = b.brandId where productId = ?", id).Scan(&prod).Error
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func EditProduct(h *gorm.DB, prod *request.Products) error {
	if err := h.Raw("UPDATE products SET productName = ?, productDescription = ? , productPrice = ?,productImage = ? ,brandId = ? WHERE productId = ?;",
		prod.ProductName,
		prod.ProductDescription,
		prod.ProductPrice,
		prod.ProductImage,
		prod.BrandId,
		prod.ProductId).Scan(&prod).Error; err != nil {
		return err
	}
	return nil
}

func DelProductColor(h *gorm.DB, id string) error {
	if err := h.Exec("DELETE FROM productcolor where productId = ?;", id).Error; err != nil {
		return err
	}
	return nil
}

func AddProductColor(h *gorm.DB, prodCl *request.OneProductColor) error {
	if err := h.Raw("INSERT INTO productcolor (productcolorId,productId,colorId) VALUES (?,?,?);",
		prodCl.ProductcolorId,
		prodCl.ProductId,
		prodCl.ColorId,
	).Scan(&prodCl).Error; err != nil {
		return err
	}
	return nil
}

func DelProduct(h *gorm.DB, id string) error {
	if err := h.Exec("DELETE FROM products where productId = ?;", id).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(h *gorm.DB, prod *request.Products) error {
	if err := h.Raw("INSERT INTO products (productId,productName,productDescription,onsaleDate,productPrice,productImage,brandId) VALUES (?,?,?,?,?,?,?);",
		prod.ProductId,
		prod.ProductName,
		prod.ProductDescription,
		time.Now().UTC().Format("2006-01-02"),
		prod.ProductPrice,
		prod.ProductImage,
		prod.BrandId,
	).Scan(&prod).Error; err != nil {
		return err
	}
	return nil
}
