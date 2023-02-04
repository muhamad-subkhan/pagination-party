package models

type Product struct {
	ID    int    `json:"id" gorm:"primaryKey:auto_increment"`
	Name  string `json:"name" gorm:"type: varchar(225)"`
	Qty   int    `json:"qty" gorm:"type: int"`
	Price int    `json:"price" gorm:"type: int"`
	Image string `json:"image" gorm:"type: varchar(225)"`
}