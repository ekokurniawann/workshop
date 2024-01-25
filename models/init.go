package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Person struct {
	gorm.Model
	Name  string `gorm:"type:varchar(300)" json:"name"`
	Age   int    `gorm:"int(5)" json:"age"`
	Email string `gorm:"type:varchar(300)" json:"email"`
}

func InitDB() {
	dsn := "host=127.0.0.1 user=echokurniawan password=12345678 dbname=belajar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err, "error connect db")
	}
	db.AutoMigrate(&Person{})

	DB = db
}
