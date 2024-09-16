/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package initialize

import (
	"log"
	"time"

	"github.com/mikietechie/gocurrenciesapi/internal/models"
)

func Init() {
	models.ConnectDb()
}

func Tear() {
	log.Println("Tearing down, will sleep for 30 seconds to allow go routines to finish")
	time.Sleep(time.Second * 30)
	models.DisonnectDb()
}
