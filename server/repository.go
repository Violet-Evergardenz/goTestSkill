package server

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	p "myapp/model/entity"
	r "myapp/model/request"
)

func (h *FuncHandler) GetAllProduct(c echo.Context) error {
	product := new([]p.Products)
	if err := h.DB.Raw("select * from products p join brands b on p.brandId = b.brandId ").Scan(&product).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	log.Println(product)
	return c.JSON(http.StatusOK, product)
}

func (h *FuncHandler) GetProduct(c echo.Context) error {

	id := c.Param("id")
	product := new(p.Products)
	// log.Println(c.Path())
	if err := h.DB.Raw("select * from products p join brands b on p.brandId = b.brandId where productId = ?", id).Scan(&product).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	log.Println(product)
	return c.JSON(http.StatusOK, product)
}

func (h *FuncHandler) EditProduct(c echo.Context) (err error) {

	product := new(r.Products)
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "something wrong")
	}

	if err := h.DB.Raw("UPDATE products SET productName = ?, productDescription = ? , productPrice = ?,productImage = ? ,brandId = ? WHERE productId = ?;",
		product.ProductName,
		product.ProductDescription,
		product.ProductPrice,
		product.ProductImage,
		product.BrandId,
		product.ProductId).Scan(&product).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, product)
}

func (h *FuncHandler) AddProduct(c echo.Context) (err error) {
	product := new(r.Products)
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "something wrong product")
	}
	log.Println(product)
	if err := h.DB.Raw("INSERT INTO products (productId,productName,productDescription,onsaleDate,productPrice,productImage,brandId) VALUES (?,?,?,?,?,?,?);",
		product.ProductId,
		product.ProductName,
		product.ProductDescription,
		time.Now().UTC().Format("2006-01-02"),
		product.ProductPrice,
		product.ProductImage,
		product.BrandId,
	).Scan(&product).Error; err != nil {
		return c.JSON(http.StatusConflict, err)
	}
	for i := 0; i < len(product.ProductColors); i++ {
		if err := h.DB.Raw("INSERT INTO productcolor (productcolorId,productId,colorId) VALUES (?,?,?);",
			product.ProductColors[i].ProductcolorId,
			product.ProductId,
			product.ProductColors[i].ColorId,
		).Scan(&product).Error; err != nil {
			return c.JSON(http.StatusConflict, err)
		}
	}
	return c.JSON(http.StatusOK, product)
}

func (h *FuncHandler) DelProduct(c echo.Context) (err error) {
	id := c.Param("id")
	// h.DB.Exec("SET FOREIGN_KEY_CHECKS=0;")
	if err := h.DB.Exec("DELETE FROM productcolor where productId = ?;", id).Error; err != nil {
		return c.JSON(http.StatusConflict, err)
	}
	if err := h.DB.Exec("DELETE FROM products where productId = ?;", id).Error; err != nil {
		return c.JSON(http.StatusConflict, err)
	}
	// h.DB.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return c.JSON(http.StatusOK, id)
}
