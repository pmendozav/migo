package main

import (
	"mio/app"
	"mio/db/gorm"
	"mio/bootstrap"
	"mio/app/models"
	"mio/app/controllers"
)

func main() {
	// init server
	app.Init()

	// init db
	gorm.Init()
	// autoDropTables()
	// autoCreateTables()
	//autoMigrateTables()

	// init redis
	// redis.Init()

	controllers.RegisterUsers()

	// app.Server.Logger.Fatal(app.Server.Start(":8080"))
	app.Server.Logger.Fatal(app.Server.StartTLS(":8080", "./config/keys/server.crt", "./config/keys/server.key"))
}

// autoMigrateTables: migrate table columns using GORM
func autoMigrateTables() {
	gorm.DBManager().AutoMigrate(&models.User{})
}

// autoCreateTables: create database tables using GORM
func autoCreateTables() {
	if !gorm.DBManager().HasTable(&models.User{}) {
		gorm.DBManager().CreateTable(&models.User{})
	}
}

// auto drop tables on dev mode
func autoDropTables() {
	if bootstrap.App.ENV == "dev" {
		gorm.DBManager().DropTableIfExists(&models.User{}, &models.User{})
	}
}