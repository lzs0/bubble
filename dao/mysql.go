package dao

import (
	"bubble/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SqlInit() {
	var err error
	DB, err = gorm.Open(mysql.Open(configs.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
