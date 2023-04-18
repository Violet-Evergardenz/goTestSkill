package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"myapp/model/shirt/entity"
	"myapp/model/shirt/request"
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
	product := new([]entity.Products)
	res, err := handle.GetAllProductHandle(h.DB, product)
	if err == nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *FuncHandler) GetProduct(c echo.Context) error {
	//check id ?
	id := c.Param("id")
	product := new(entity.Products)
	res, err := handle.GetProductHandle(h.DB, product, id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if res.ProductId == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *FuncHandler) EditProduct(c echo.Context) (err error) {
	product := new(request.Products)
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "something wrong")
	}
	// log.Println(product.ProductColors[0])
	res, err := handle.EditProductHandle(h.DB, product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *FuncHandler) AddProduct(c echo.Context) (err error) {
	product := new(request.Products)
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "something wrong product")
	}
	if err := handle.AddProductHandle(h.DB, product); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	log.Println(product)
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
