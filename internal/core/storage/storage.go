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
		logger.ErrorLog.Fatalln("Problem setting SqlStorer instance")
	}

	return
}

func connection() (d *sql.DB, err error) {
	PATH := os.Getenv("DB_ADDR")
	if PATH == "" {
		PATH = "INVALID_ENV_ADDR"
	}

	d, err = sql.Open("mysql", PATH)
	if err != nil {
		logger.ErrorLog.Println("Failed to Open connection with PATH: ", PATH)
		return nil, sql.ErrConnDone
	}
	logger.InfoLog.Println("Established Connection to database")
	d.SetConnMaxLifetime(time.Minute * 3)
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(10)

	return
}
