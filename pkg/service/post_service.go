package service

import (
	"fmt"
	"log"

	"tidy/pkg/model"
	"tidy/pkg/repository"
)

type PostServ struct {
	storage repository.PostStorage
}

// NewService ...
func NewPostService(storage repository.PostStorage) *PostServ {
	return &PostServ{
		storage: storage,
	}
}
func (s *PostServ) GetAll() ([]model.Post, error) {
	p, err := s.storage.GetAll()
	if err != nil {
		log.Printf("ERROR post  PostServ GetAll method :--> %v\n", err)
		return p, err
	}
	log.Println("GetAll method post  PostServ")

	return p, nil
}
func (s *PostServ) GetSearch(str string) ([]model.Post, error) {
	var p []model.Post
	var err error
	if len(str) > 0 {
		p, err = s.storage.GetSearch(str)
		if err != nil {
			log.Printf("ERROR post  PostServ GetAll method :--> %v\n", err)
			return p, err
		}
	} else {
		p, err = s.storage.GetAll()
		if err != nil {
			log.Printf("ERROR post  PostServ GetAll method :--> %v\n", err)
			return p, err
		}
	}

	return p, nil
}
func (s *PostServ) PostPage(id int) ([]model.Post, error) {

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
func (s *PostServ) GetCategory() ([]model.Category, error) {
	p, err := s.storage.GetCategory()
	if err != nil {
		log.Printf("ERROR post  PostServ GetCategory method :--> %v\n", err)
		return p, err
	}
	log.Println("GetAll method post  PostServ")

	return p, nil
}
func (s *PostServ) SortedByCategory(t string) ([]model.Post, error) {
	M, err := s.storage.SortedCategory(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}

func (s *PostServ) LikedPosts(t int) ([]model.Post, error) {
	M, err := s.storage.LikedPosts(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}
func (s *PostServ) CreatedPosts(t int) ([]model.Post, error) {
	M, err := s.storage.CreatedPosts(t)
	if err != nil {
		fmt.Println(err)
		return M, err
	}
	return M, nil
}
func (s *PostServ) CreatePost(m model.Post, cat []string, id int) model.Post {
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
