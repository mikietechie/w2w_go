/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package config

import (
	"context"

	"github.com/joho/godotenv"
)

var _ = godotenv.Overload()
var CTX = context.Background()

/* System */
var SYS_NAME = GetEnvOrDef("SYS_NAME", "W2W GO API")
var SECRET_KEY = GetEnvOrDef("SECRET_KEY", "MikeAndGo")
var ENV = GetEnvOrDef("ENV", "PROD")
var DEV = ENV == "DEV"
var SERVER_ADDRESS = GetEnvOrDef("SERVER_ADDRESS", "0.0.0.0:8000")

/* POSTGRES MONGO REDIS */
var DATABASE_CONNECTION = GetEnvOrDef(
	"POSTGRES_CONNECTION",
	"postgres://pg:pass@localhost:5432/db",
)
