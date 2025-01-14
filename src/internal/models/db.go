/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package models

import (
	"log"

	"github.com/mikietechie/gocurrenciesapi/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	Db, err = gorm.Open(postgres.Open(config.DATABASE_CONNECTION), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database\n", err)
	}
	// Db.Logger.LogMode(logger.Silent)
	Db.AutoMigrate(&RoomModel{}, &UserModel{}, &RoomUserModel{})
}

func DisonnectDb() {
	db, err := Db.DB()
	if err != nil {
		log.Println("Failed to get DB to close it, most probably it is already closed")
	}
	err = db.Close()
	if err != nil {
		log.Println("Failed to get DB to close it, most probably it is already closed")
	} else {
		log.Println("Db Closed")
	}
}
