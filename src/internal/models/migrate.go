/*
Date Created		14 September 2024
Author				Mike Z
Email				mzinyoni7@outlook.com
Website				https://mikeio.web.app
Status				Looking for a job!
Description			Users and rooms simulation
*/

package models

func Migrate() {
	Db.AutoMigrate(&RoomModel{}, &UserModel{}, &RoomUserModel{})
}
