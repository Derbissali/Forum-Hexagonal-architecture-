package user

import "tezt/hexagonal/internal/model"

type UserStorage interface {
	Create(m *model.User) (model.User, error)
	Check(n, p string) bool
	SelectUserID(m string) (string, error)
}
