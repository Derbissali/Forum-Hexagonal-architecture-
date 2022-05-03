package repository

import (
	"database/sql"
	"tidy/pkg/model"
)

type PostStorage interface {
	GetAll() ([]model.Post, error)
	GetSearch(str string) ([]model.Post, error)
	GetCategory() ([]model.Category, error)
	SortedCategory(t string) ([]model.Post, error)
	LikedPosts(t int) ([]model.Post, error)
	CreatedPosts(t int) ([]model.Post, error)
	CreatePost(t model.Post, s []string, ID int) error
	CountPost() (int, error)
	SinglePost(id int) ([]model.Post, error)
}
type UserStorage interface {
	Create(m *model.User) (model.User, error)
	Check(n, p string) bool
	SelectUserID(m string) (string, error)
}

type SessionStorage interface {
	Create(uuid, id string) error
	ReadByUUID(uuid string) (model.User, error)
	Delete(n string) error
}
type LikeStorage interface {
	SetLike(idPost string, n int)
	SetDislike(idPost string, n int)
	UpdateLike(idPost string, n int)
	UpdateDislike(idPost string, n int)
	PostDislike(n int, idPost string) int
	PostLike(n int, idPost string) int
	DeleteLikeNDis(idPost string, n int)
	UpdateLikeCount(idPost string)
	UpdateDislikeCount(idPost string)

	SetCommentLike(idPost, idComment string, n int)
	SetCommentDislike(idPost, idComment string, n int)
	UpdateCommentLike(idPost, idComment string, n int)
	UpdateCommentDislike(idPost, idComment string, n int)
	CommentDislike(n int, idPost, idComment string) int
	CommentLike(n int, idPost, idComment string) int
	DeleteCommentLikeNDis(idPost, idComment string, n int)
	UpdateCommentLikeCount(idPost, idComment string)
	UpdateCommentDislikeCount(idPost, idComment string)
}
type CommentStorage interface {
	AddComment(comment string, id, n int) error
}
type Repository struct {
	PostStorage
	UserStorage
	SessionStorage
	LikeStorage
	CommentStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		PostStorage:    NewPostStorage(db),
		SessionStorage: NewSessionStorage(db),
		LikeStorage:    NewLikeStorage(db),
		CommentStorage: NewCommentStorage(db),
		UserStorage:    NewUserStorage(db),
	}
}
