package main

import (
	"log"
	db "myapp/server"
	"net/http"

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

	product := e.Group("/product")
	product.GET("/getProduct", h.GetAllProduct)
	product.GET("/getProduct/:id", h.GetProduct)
	product.PUT("/editProduct", h.EditProduct)
	product.POST("/addProduct", h.AddProduct)
	product.DELETE("/delProduct/:id", h.DelProduct)

	chatchak := e.Group("/chatchak")
	chatchak.GET("/join/:id", h.JoinChatRoom)

	e.Logger.Fatal(e.Start(":1323"))
}
