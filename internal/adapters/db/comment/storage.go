package comment

import (
	"database/sql"
	"fmt"
	"tezt/hexagonal/internal/domain/comment"
)

type commentStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) comment.CommentStorage {
	return &commentStorage{
		db: db,
	}
}

func (s *commentStorage) AddComment(comment string, id, n int) error {
	_, err := s.db.Exec(`INSERT INTO comment (body, post_id, user_id, likes, dislikes) VALUES (?, ?, ?,?,?)`, comment, id, n, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
