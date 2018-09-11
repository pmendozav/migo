package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"mio/bootstrap"
)

var (
	db *gorm.DB
)

// initialize db
func Init() {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")

	switch adapter {
		case "mysql":
			mysqlConn()
			break
		case "postgres":
			//postgresConn()
			break
	}
}

// mysqlConn: setup mysql database connection using the configuration from database.yaml
func mysqlConn() {
	var (
		connectionString string
		err              error
	)
	
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"))
	
	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

/*
 * Gorm: return GORM's postgres database connection instance.
 */
 func DBManager() *gorm.DB {
	return db
}
