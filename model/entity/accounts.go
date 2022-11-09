package entity

type Accounts struct {
	AccountId   int    `json:"accountId" gorm:"column:accountId"`
	AccountName string `json:"accountName" gorm:"column:accountName"`
	UserName    string `json:"userName" gorm:"column:userName"`
	Password    string `json:"password" gorm:"column:password"`
	AccountRole string `json:"accountRole" gorm:"column:accountRole"`
}
