package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var DB_PASSWORD = "root"
	var DB_HOST = "localhost"
	var DB_PORT = "3306"
	var DB_NAME = "party"
	var err error
	// dsn := "root:@tcp(localhost:3306)/waysfood?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("mysql error")
	}

	fmt.Println("Connected to Database")
}
