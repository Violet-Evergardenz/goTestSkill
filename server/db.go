package server

import (
	"fmt"

	"log"
	"net/http"
	"os"
	"time"

	p "myapp/model/entity"
	r "myapp/model/request"
	"myapp/server/service/handle"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type FuncHandler struct {
	DB *gorm.DB
}

func (h *FuncHandler) Initialize() {
	dns := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local`, os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	h.DB = db
}

func (h *FuncHandler) GetAllProduct(c echo.Context) error {
	res, err := handle.GetAllProductHandle(c, h.DB)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
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
	log.Println(product.ProductColors[0])
	if err := h.DB.Raw("UPDATE products SET productName = ?, productDescription = ? , productPrice = ?,productImage = ? ,brandId = ? WHERE productId = ?;",
		product.ProductName,
		product.ProductDescription,
		product.ProductPrice,
		product.ProductImage,
		product.BrandId,
		product.ProductId).Scan(&product).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := h.DB.Exec("DELETE FROM productcolor where productId = ?;", product.ProductId).Error; err != nil {
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
