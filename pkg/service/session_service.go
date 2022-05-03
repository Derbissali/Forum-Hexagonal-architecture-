package service

import (
	"fmt"

	"tidy/pkg/model"
	"tidy/pkg/repository"
)

type SessionServ struct {
	storage repository.SessionStorage
}

func NewSessionService(storage repository.SessionStorage) *SessionServ {
	return &SessionServ{
		storage: storage,
	}
}

func (s *SessionServ) Create(uuid, id string) error {

	err := s.storage.Create(uuid, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *SessionServ) Delete(id string) error {
	fmt.Println(id, "IDIWKA")
	err := s.storage.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *SessionServ) ReadByUUID(uuid string) (model.User, error) {
	m, err := s.storage.ReadByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return m, err
	}
	return m, nil
}
func (s *SessionServ) Check(uuid string) error {
	return nil
}
