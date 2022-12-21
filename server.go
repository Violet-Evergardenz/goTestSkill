package main

import (
	"log"
	db "myapp/server"
	"net/http"

	// c "myapp/controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	h := db.FuncHandler{}
	h.Initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// http://localhost:1323/getProduct/1
	product := e.Group("/product")
	product.GET("/getProduct", h.GetAllProduct)
	product.GET("/getProduct/:id", h.GetProduct)
	product.PUT("/editProduct", h.EditProduct)
	product.POST("/addProduct", h.AddProduct)
	product.DELETE("/delProduct/:id", h.DelProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
