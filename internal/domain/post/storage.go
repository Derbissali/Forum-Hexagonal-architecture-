package post

import "tezt/hexagonal/internal/model"

type PostStorage interface {
	GetAll() ([]model.Post, error)
	GetCategory() ([]model.Category, error)
	SortedCategory(t string) ([]model.Post, error)
	LikedPosts(t int) ([]model.Post, error)
	CreatedPosts(t int) ([]model.Post, error)
	CreatePost(t model.Post, s []string, ID int) error
	CountPost() (int, error)
	SinglePost(id int) ([]model.Post, error)
}
