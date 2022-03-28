package comment

import (
	"fmt"
	"strconv"
	"tezt/hexagonal/internal/adapters/api"
)

type service struct {
	storage CommentStorage
}

func NewService(storage CommentStorage) api.CommentService {
	return &service{
		storage: storage,
	}
}

func (s *service) AddComment(comment string, id string, n int) error {
	var err error
	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(comment) > 0 && len(comment) < 140 {
		s.storage.AddComment(comment, i, n)
	} else {
		return err
	}
	return nil
}
