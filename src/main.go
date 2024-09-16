/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package main

import (
	"github.com/mikietechie/gocurrenciesapi/internal/initialize"
	"github.com/mikietechie/gocurrenciesapi/internal/server"
)

// @title                       W2W-GO
// @version                     0.1
// @description                 W2W-GO
// @contact.name                Mike Zinyoni
// @contact.email               mzinyoni7
func main() {
	initialize.Init()
	defer initialize.Tear()
	server.RunServer()
}
