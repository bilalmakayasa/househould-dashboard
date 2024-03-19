package database

import (
	"household-dashboard/src/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", config.DbConfig())
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
