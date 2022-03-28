package dbase

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CheckDB() *sql.DB {

	var d Database
	var err error

	d.db, err = sql.Open("sqlite3", "dbase/forum.db")

	if err != nil {

		log.Fatalf("this error is in dbase/open() %v", err)
	}
	return d.db
}
