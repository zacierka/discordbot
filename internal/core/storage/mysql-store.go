package storage

import (
	"database/sql"
	"discordbot/internal/logger"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBInstance *sql.DB

func SetupStorage() error {
	PATH := os.Getenv("DB_ADDR")
	if PATH == "" {
		PATH = "INVALID-ADDRESS"
	}

	db, err := sql.Open("mysql", PATH)
	if err != nil {
		logger.ErrorLog.Println("Could not Connect to Database on ADDR: ", PATH)
		return sql.ErrConnDone
	}
	logger.InfoLog.Println("Established Connection to database")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DBInstance = db

	return nil
}

func PingDB() bool {
	err := DBInstance.Ping()
	if err != nil {
		return false
	}
	logger.InfoLog.Println("ERROR IS ", err)
	return true
}
