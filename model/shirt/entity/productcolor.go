package entity

type Productcolor struct {
	ProductcolorId int `json:"productcolorId" gorm:"column:productcolorId"`
	ProductId      int `json:"productId" gorm:"column:productId"`
	ColorId        int `json:"colorId" gorm:"column:colorId"`
}
