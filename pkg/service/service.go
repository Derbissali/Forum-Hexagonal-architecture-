package service

import (
	"tidy/pkg/model"
	"tidy/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type PostService interface {
	GetAll() ([]model.Post, error)
	GetSearch(str string) ([]model.Post, error)
	GetCategory() ([]model.Category, error)
	SortedByCategory(t string) ([]model.Post, error)
	LikedPosts(t int) ([]model.Post, error)
	CreatedPosts(t int) ([]model.Post, error)
	CreatePost(m model.Post, cat []string, id int) model.Post
	PostPage(id int) ([]model.Post, error)
}
type LikeService interface {
	PostLike(l, d, idPost string, n int) error
	CommentLike(l, d, idPost, idComment string, n int) error
	SetLikeDislike(l, d, idPost string, n int) error
	SetLikeDislikeComment(l, d, idPost, idComment string, n int) error
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

type Service struct {
	UserService
	PostService
	CommentService
	SessionService
	LikeService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService:    NewUserService(repos.UserStorage),
		PostService:    NewPostService(repos.PostStorage),
		CommentService: NewCommentService(repos.CommentStorage),
		SessionService: NewSessionService(repos.SessionStorage),
		LikeService:    NewLikeService(repos.LikeStorage),
	}
}
