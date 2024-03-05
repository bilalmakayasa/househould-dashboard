package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func ConnectDB() {
	dbinfo := fmt.Sprintf("sslmode=disable dbname=%s host=%s port=%s user=%s password=%s", viper.GetString("database.name"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.user"), viper.GetString("database.password"))

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
