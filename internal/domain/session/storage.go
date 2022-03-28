package session

import "tezt/hexagonal/internal/model"

type SessionStorage interface {
	Create(uuid, id string) error
	ReadByUUID(uuid string) (model.User, error)
	Delete(n string) error
}
