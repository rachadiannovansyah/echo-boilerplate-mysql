package database

import (
	"database/sql"
	"learn-go-echo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connection failure!")
	}

	err = db.Ping()
	if err != nil {
		panic("Invalid!")
	}
}

func CreateConnection() *sql.DB {
	return db
}
