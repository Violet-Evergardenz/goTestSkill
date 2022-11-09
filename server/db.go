package server

import (
	"fmt"

	"os"

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
