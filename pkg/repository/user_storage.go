package repository

import (
	"database/sql"
	"fmt"
	"tidy/pkg/model"

	"golang.org/x/crypto/bcrypt"
)

type SqlUserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *SqlUserStorage {
	return &SqlUserStorage{
		db: db,
	}
}

func (c *SqlUserStorage) Create(m *model.User) (model.User, error) {
	checkUniq := checkUniq(c, m.Name, m.Email)
	if !checkUniq {
		m.ErrorEm = true
		return *m, nil
	}
	_, err := c.db.Exec(`INSERT INTO user (name, email, password) VALUES (?, ?, ?)`, m.Name, m.Email, m.Password)
	if err != nil {
		fmt.Println(err)
		return *m, err
	}
	return *m, nil
}

func checkUniq(c *SqlUserStorage, name, Email string) bool {
	notUn := 0
	notUniq := c.db.QueryRow((`SELECT user.id FROM user WHERE user.email=?`), Email)
	notUniq.Scan(&notUn)
	if notUn != 0 {
		fmt.Println(notUn)
		return false
	}

	return true
}

func (c *SqlUserStorage) Check(n, p string) bool {
	result := c.db.QueryRow(`SELECT "password" from "user" WHERE email=$1`, n)
	fmt.Println(n, p)
	ourPerson := model.User{}
	err := result.Scan(&ourPerson.Password)
	if err == sql.ErrNoRows {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		fmt.Println("wrong login")
		return false

	}
	//	fmt.Println(ourPerson.Password, creds.Password)
	if err = bcrypt.CompareHashAndPassword([]byte(ourPerson.Password), []byte(p)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("wrong password")
		return false
	}
	return true
}

func (c *SqlUserStorage) SelectUserID(name string) (string, error) {
	a := ""
	row := c.db.QueryRow((`SELECT user.id FROM user WHERE user.email = ?`), name)
	e := row.Scan(&a)
	if e != nil {
		fmt.Println(e)
		return "", e
	}
	return a, nil
}
