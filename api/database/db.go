package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hardware-store/api/database/config"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.BuildDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
