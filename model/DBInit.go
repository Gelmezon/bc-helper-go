package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DbClient *gorm.DB

func DBInit() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/bc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Default().Println(err.Error())
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&BcWear{})

	DbClient = db

}
