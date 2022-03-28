package user

import (
	"log"
	"tezt/hexagonal/internal/adapters/api"
	"tezt/hexagonal/internal/adapters/utils"
	"tezt/hexagonal/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	storage UserStorage
}

// NewService ...
func NewService(storage UserStorage) api.UserService {
	return &service{storage: storage}
}

func (c *service) Create(m *model.User) (model.User, error) {
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

func (c *service) Check(m *model.User) (model.User, error) {

	if !c.storage.Check(m.Email, m.Password) {
		m.ErrorEm = true
		return *m, nil
	}
	return *m, nil
}
func (c *service) GetIDbyName(m string) (string, error) {
	id, err := c.storage.SelectUserID(m)
	if err != nil {
		return "", err
	}

	return id, nil
}
