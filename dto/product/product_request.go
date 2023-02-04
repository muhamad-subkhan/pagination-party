package productdto

type ProductRequest struct {
	Name  string `json:"name" gorm:"type: varchar(225)"`
	Qty   int    `json:"qty" gorm:"type: int"`
	Price int    `json:"price" gorm:"type: int"`
	Image string `json:"image" gorm:"type: varchar(225)"`
}

type UpdateProductRequest struct {
	Name  string `json:"name" gorm:"type: varchar(225)"`
	Qty   int    `json:"qty" gorm:"type: int"`
	Price int    `json:"price" gorm:"type: int"`
	Image string `json:"image" gorm:"type: varchar(225)"`
}
