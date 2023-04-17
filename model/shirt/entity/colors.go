package entity

type Colors struct {
	ColorId    int    `json:"colorId" gorm:"column:colorId"`
	ColorName  string `json:"colorName" gorm:"column:colorName"`
	ColorValue string `json:"colorValue" gorm:"column:colorValue"`
}
