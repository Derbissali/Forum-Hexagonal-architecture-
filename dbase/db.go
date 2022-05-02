package dbase

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CheckDB() *sql.DB {

	_, err := os.Stat("dbase/database-sqlite.db")
	if os.IsNotExist(err) {
		createFile()
	}
	var d Database
	d.open("dbase/database-sqlite.db")
	d.createTable()
	return d.db
}
func createFile() {
	file, err := os.Create("dbase/database-sqlite.db")
	if err != nil {
		log.Fatalf("file doesn't create %v", err)
	}
	defer file.Close()
}

func (d *Database) open(file string) {
	var err error
	d.db, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("this error is in dbase/open() %v", err)
	}
}

func (d *Database) createTable() {
	_, err := d.db.Exec(`CREATE TABLE IF NOT EXISTS "user" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		"email"	TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE user")
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "post" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		"body"	TEXT NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"likes"	INTEGER,
		"dislikes"	INTEGER,
		"image"	TEXT,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE posts")
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "category" (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE categories")
	}
	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "category_post" (
		"id"	INTEGER NOT NULL,
		"category_id"	INTEGER NOT NULL,
		"post_id"	INTEGER NOT NULL,
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("category_id") REFERENCES "category"("id")
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE post_category")
	}

	_, err = d.db.Exec(`INSERT INTO category (name)
		VALUES('Dormitory'),
			('University Questions'),
			('Software Engineers'),
			('Questions to teachers'),
			('Entertainment'),
			('Practice skills'),
			('Hubby');`)
	if err != nil {
		// log.Printf("%v\n", err)
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "comment" (
		"id"	INTEGER NOT NULL,
		"body"	TEXT NOT NULL,
		"post_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"likes"	INTEGER,
		"dislikes"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE comments")
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "likeNdis" (
		"id"	INTEGER NOT NULL,
		"like"	INTEGER,
		"dislike"	INTEGER,
		"post_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		FOREIGN KEY("user_id") REFERENCES "user"("id"),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE likes")
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "comment_like_dislike" (
		"id"	INTEGER NOT NULL,
		"like"	INTEGER,
		"dislike"	INTEGER,
		"comment_id"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL,
		"post_id"	INTEGER NOT NULL,
		PRIMARY KEY("id" AUTOINCREMENT),
		FOREIGN KEY("comment_id") REFERENCES "comment"("id"),
		FOREIGN KEY("post_id") REFERENCES "post"("id"),
		FOREIGN KEY("user_id") REFERENCES "user"("id")
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE commentslikes")
	}

	_, err = d.db.Exec(`CREATE TABLE IF NOT EXISTS "session" (
		"uuid"	INTEGER NOT NULL,
		"user_id"	INTEGER NOT NULL UNIQUE
	);`)
	if err != nil {
		log.Println("CAN NOT CREATE TABLE sessions")
	}
}
