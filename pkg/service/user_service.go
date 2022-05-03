package service

import (
	"log"
	"tidy/pkg/model"
	"tidy/pkg/repository"
	"tidy/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserServ struct {
	storage repository.UserStorage
}

// NewService ...
func NewUserService(storage repository.UserStorage) *UserServ {
	return &UserServ{storage: storage}
}

func (c *UserServ) Create(m *model.User) (model.User, error) {
	empty := utils.CheckEmpty(m.Name, m.Email, m.Password)
	if !empty {
		m.ErrorEmpty = true
		return *m, nil
	}
	CheckEmail := utils.CheckEmail(m.Email)
	if !CheckEmail {
		m.ErrorE = true
		return *m, nil
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), 8)
	m.Password = string(hashedPassword)
	if err != nil {
		log.Println("Error server in user creation")
		return *m, err
	}
	user, err := c.storage.Create(m)
	if err != nil {
		log.Println("Error server in user creation")
		return *m, err
	}
	return user, nil
}

func (c *UserServ) Check(m *model.User) (model.User, error) {

	if !c.storage.Check(m.Email, m.Password) {
		m.ErrorEm = true
		return *m, nil
	}
	return *m, nil
}
func (c *UserServ) GetIDbyName(m string) (string, error) {
	id, err := c.storage.SelectUserID(m)
	if err != nil {
		return "", err
	}

	return id, nil
}
