package post

import (
	"fmt"
	"log"
	"tezt/hexagonal/internal/adapters/api"
	"tezt/hexagonal/internal/domain/like"
	"tezt/hexagonal/internal/model"
)

type service struct {
	storage     PostStorage
	likeService like.LikeService
}

// NewService ...
func NewService(storage PostStorage, likeService like.LikeService) api.PostService {
	return &service{
		storage:     storage,
		likeService: likeService,
	}
}
func (s *service) GetAll() ([]model.Post, error) {
	p, err := s.storage.GetAll()
	if err != nil {
		log.Printf("ERROR post service GetAll method :--> %v\n", err)
		return p, err
	}
	log.Println("GetAll method post service")

	return p, nil
}
func (s *service) GetSearch(str string) ([]model.Post, error) {
	var p []model.Post
	var err error
	if len(str) > 0 {
		p, err = s.storage.GetSearch(str)
		if err != nil {
			log.Printf("ERROR post service GetAll method :--> %v\n", err)
			return p, err
		}
	} else {
		p, err = s.storage.GetAll()
		if err != nil {
			log.Printf("ERROR post service GetAll method :--> %v\n", err)
			return p, err
		}
	}

	return p, nil
}
func (s *service) PostPage(id int) ([]model.Post, error) {

	posts, err := s.storage.CountPost()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if id < 1 || id > posts {
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil, nil
	}
	M, err := s.storage.SinglePost(id)
	if err != nil {
		return nil, err
	}
	return M, err
}
func (s *service) GetCategory() ([]model.Category, error) {
	p, err := s.storage.GetCategory()
	if err != nil {
		log.Printf("ERROR post service GetCategory method :--> %v\n", err)
		return p, err
	}
	log.Println("GetAll method post service")

	return p, nil
}
func (s *service) SortedByCategory(t string) ([]model.Post, error) {
	M, err := s.storage.SortedCategory(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}

func (s *service) LikedPosts(t int) ([]model.Post, error) {
	M, err := s.storage.LikedPosts(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}
func (s *service) CreatedPosts(t int) ([]model.Post, error) {
	M, err := s.storage.CreatedPosts(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}
func (s *service) CreatePost(m model.Post, cat []string, id int) model.Post {
	if len(m.Name) > 50 || len(m.Body) > 2000 {
		m.TitBodOver = true
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return m
	}
	if len(m.Name) == 0 || len(m.Body) == 0 {
		m.TitBodNull = true
		return m
	}

	if len(cat) <= 0 {
		m.CategoryNull = true
		return m
	}

	err := s.storage.CreatePost(m, cat, id)
	if err != nil {
		m.CategoryNull = true
		return m
	}
	return m
}
func (s *service) PostLike(l, d, idPost string, n int) error {
	if err := s.likeService.SetLikeDislike(l, d, idPost, n); err != nil {
		log.Printf("ERROR post service Check Post Like method:----> %v\n", err)
		return err
	}
	return nil
}
func (s *service) CommentLike(l, d, idPost, idComment string, n int) error {
	if err := s.likeService.SetLikeDislikeComment(l, d, idPost, idComment, n); err != nil {
		log.Printf("ERROR post service Check Post Like method:----> %v\n", err)
		return err
	}
	return nil
}
