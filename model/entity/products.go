package entity

// "gorm.io/gorm"

type Products struct {
	// gorm.Model
	ProductId          int     `json:"productId" gorm:"column:productId"`
	ProductName        string  `json:"productName" gorm:"column:productName"`
	ProductDescription string  `json:"productDescription" gorm:"column:productDescription"`
	OnsaleDate         string  `json:"onsaleDate" gorm:"column:onsaleDate"`
	ProductPrice       float64 `json:"productPrice" gorm:"column:productPrice"`
	ProductImage       string  `json:"productImage" gorm:"column:productImage"`
	BrandId            int     `json:"BrandId" gorm:"column:brandId"`
	// Brands             []*Brands `gorm:"foreignKey:BrandId; references:BrandId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BrandName string `json:"BrandName" gorm:"column:brandName"`
}
