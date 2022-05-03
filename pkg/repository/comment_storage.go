package repository

import (
	"database/sql"
	"fmt"
)

type SqlCommentStorage struct {
	db *sql.DB
}

func NewCommentStorage(db *sql.DB) *SqlCommentStorage {
	return &SqlCommentStorage{
		db: db,
	}
}

func (s *SqlCommentStorage) AddComment(comment string, id, n int) error {
	_, err := s.db.Exec(`INSERT INTO comment (body, post_id, user_id, likes, dislikes) VALUES (?, ?, ?,?,?)`, comment, id, n, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
