package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Initdb() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/powergophers?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return (err)
	}

	return nil
}
