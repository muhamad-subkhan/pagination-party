package dto

type RegisterResponse struct {
	Token string `json:"token" gorm:"type: varchar(225)"`
}

type LoginResponse struct {
	Token string `json:"token" gorm:"type: varchar(225)"`
}