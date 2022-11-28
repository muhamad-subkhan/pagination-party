package models

type Profile struct {
	ID       int    `json:"id" gorm:"primaryKey:auto_increment"`
	FullName string `json:"FullName" gorm:"type: varchar(225)"`
	Email    string `json:"email" gorm:"type: varchar(225)"`
	Password string `json:"password" gorm:"type: varchar(225)"`
	Image    string `json:"image" gorm:"type: varchar(225)"`
}
