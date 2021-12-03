package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbDsn string) {

	db, err := gorm.Open(mysql.Open(dbDsn))
	if err != nil {
		panic(err.Error())
		return
	}

	db.AutoMigrate(&Book{})

	DB = db
}
