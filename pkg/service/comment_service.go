package service

import (
	"fmt"
	"strconv"
	"tidy/pkg/repository"
)

type CommentServ struct {
	storage repository.CommentStorage
}

func NewCommentService(storage repository.CommentStorage) *CommentServ {
	return &CommentServ{
		storage: storage,
	}
}

func (s *CommentServ) AddComment(comment string, id string, n int) error {
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
