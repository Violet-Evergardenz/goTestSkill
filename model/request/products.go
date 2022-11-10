package request

type Products struct {
	ProductId          int               `json:"productId" gorm:"column:productId"`
	ProductName        string            `json:"productName" gorm:"column:productName"`
	ProductDescription string            `json:"productDescription" gorm:"column:productDescription"`
	OnsaleDate         string            `json:"onsaleDate" gorm:"column:onsaleDate"`
	ProductPrice       float64           `json:"productPrice" gorm:"column:productPrice"`
	ProductImage       string            `json:"productImage" gorm:"column:productImage"`
	BrandId            int               `json:"brandId" gorm:"column:brandId"`
	ProductColors      []OneProductColor `json:"productColor"`
}

type OneProductColor struct {
	ProductcolorId int `json:"productcolorId" gorm:"column:productcolorId"`
	ProductId      int `json:"productId" gorm:"column:productId"`
	ColorId        int `json:"colorId" gorm:"column:colorId"`
}
