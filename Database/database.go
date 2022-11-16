package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() error {
	db, err := gorm.Open(mysql.Open("backendGolang"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate()
	return nil
}
