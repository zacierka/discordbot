package storage

import (
	"database/sql"
	"discordbot/internal/logger"
	"os"
	"time"
)

type Storer interface {
	GetVersion() string
}

type SqlStorer struct {
	instance *sql.DB
}

func New() (s *SqlStorer, err error) {
	s = &SqlStorer{}
	s.instance, err = connection()
	if err != nil {
		logger.ErrorLog.Fatalln("Cannot get database connection")
	}

	return
}

func connection() (d *sql.DB, err error) {
	PATH := os.Getenv("DB_ADDR")
	if PATH == "" {
		PATH = "INVALID-ADDRESS"
	}

	logger.InfoLog.Println("Path is ", PATH)

	db, err := sql.Open("mysql", PATH)
	if err != nil {
		logger.ErrorLog.Println("Could not Connect to Database on ADDR: ", PATH)
		return nil, sql.ErrConnDone
	}
	logger.InfoLog.Println("Established Connection to database")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d = db

	return d, nil
}
