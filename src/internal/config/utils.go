/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package config

import "os"

func GetEnvOrDef(key, def string) string {
	env := os.Getenv(key)
	if len(env) != 0 {
		return env
	}
	return def
}
