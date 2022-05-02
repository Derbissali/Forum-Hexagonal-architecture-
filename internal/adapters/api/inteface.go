package api

import (
	"net/http"
	"tezt/hexagonal/internal/model"
)

type Handler interface {
	Register(router *http.ServeMux)
}
type PostService interface {
	GetAll() ([]model.Post, error)
	GetSearch(str string) ([]model.Post, error)
	GetCategory() ([]model.Category, error)
	SortedByCategory(t string) ([]model.Post, error)
	LikedPosts(t int) ([]model.Post, error)
	CreatedPosts(t int) ([]model.Post, error)
	CreatePost(m model.Post, cat []string, id int) model.Post
	PostPage(id int) ([]model.Post, error)
	PostLike(l, d, idPost string, n int) error
	CommentLike(l, d, idPost, idComment string, n int) error
}
type UserService interface {
	Create(m *model.User) (model.User, error)
	Check(m *model.User) (model.User, error)
	GetIDbyName(email string) (string, error)
}
type SessionService interface {
	Create(uuid, id string) error
	Delete(id string) error
	ReadByUUID(uuid string) (model.User, error)
}
type CommentService interface {
	AddComment(comment string, id string, n int) error
}
