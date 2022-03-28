package session

import (
	"fmt"
	"tezt/hexagonal/internal/adapters/api"
	"tezt/hexagonal/internal/model"
)

type service struct {
	storage SessionStorage
}

func NewService(storage SessionStorage) api.SessionService {
	return &service{
		storage: storage,
	}
}

func (s *service) Create(uuid, id string) error {

	err := s.storage.Create(uuid, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *service) Delete(id string) error {
	fmt.Println(id, "IDIWKA")
	err := s.storage.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *service) ReadByUUID(uuid string) (model.User, error) {
	m, err := s.storage.ReadByUUID(uuid)
	if err != nil {
		fmt.Println(err)
		return m, err
	}
	return m, nil
}
func (s *service) Check(uuid string) error {
	return nil
}
