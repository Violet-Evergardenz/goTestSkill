package entity

type Brands struct {
	// gorm.Model
	BrandId   int    `json:"brandId" gorm:"column:brandId"`
	BrandName string `json:"brandName" gorm:"column:brandName"`
}
