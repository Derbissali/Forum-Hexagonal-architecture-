package repository

import (
	"database/sql"
	"fmt"
)

type SqlLikeStorage struct {
	db *sql.DB
}

func NewLikeStorage(db *sql.DB) *SqlLikeStorage {
	return &SqlLikeStorage{
		db: db,
	}
}

func (s *SqlLikeStorage) SetLike(idPost string, n int) {
	_, err := s.db.Exec(`INSERT INTO likeNdis (like, post_id, user_id) VALUES (?, ?, ?)`, 1, idPost, n)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *SqlLikeStorage) SetDislike(idPost string, n int) {
	_, err := s.db.Exec(`INSERT INTO likeNdis (dislike, post_id, user_id) VALUES (?, ?, ?)`, 1, idPost, n)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *SqlLikeStorage) UpdateLike(idPost string, n int) {
	_, err := s.db.Exec(`UPDATE likeNdis SET like=1, dislike=NULL WHERE post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) UpdateDislike(idPost string, n int) {
	_, err := s.db.Exec(`UPDATE likeNdis SET like=NULL, dislike=1 WHERE post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *SqlLikeStorage) PostDislike(n int, idPost string) int {
	a := 0
	stmt := s.db.QueryRow(`SELECT "dislike" FROM "likeNdis" WHERE user_id=? AND post_id=?`, n, idPost)
	stmt.Scan(&a)
	return a
}
func (s *SqlLikeStorage) PostLike(n int, idPost string) int {
	a := 0
	stmtl := s.db.QueryRow(`SELECT "like" FROM "likeNdis" WHERE user_id=? AND post_id=?`, n, idPost)
	stmtl.Scan(&a)
	return a
}
func (s *SqlLikeStorage) DeleteLikeNDis(idPost string, n int) {
	_, err := s.db.Exec(`DELETE FROM likeNdis where post_id = ? AND user_id=?`, idPost, n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) UpdateLikeCount(idPost string) {
	a := 0
	row := s.db.QueryRow(`SELECT COUNT(like) FROM likeNdis WHERE post_id=?`, idPost)
	e := row.Scan(&a)
	if e != nil {
		return
	}
	_, err := s.db.Exec(`UPDATE post SET likes=? WHERE id=?`, a, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *SqlLikeStorage) UpdateDislikeCount(idPost string) {
	a := 0
	row := s.db.QueryRow(`SELECT COUNT(dislike) FROM likeNdis WHERE post_id=?`, idPost)
	e := row.Scan(&a)
	if e != nil {
		return
	}
	_, err := s.db.Exec(`UPDATE post SET dislikes=? WHERE id=?`, a, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *SqlLikeStorage) SetCommentLike(idPost, idComment string, n int) {
	_, err := s.db.Exec(`INSERT INTO comment_like_dislike (like, post_id, user_id, comment_id) VALUES (?, ?, ?, ?)`, 1, idPost, n, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) SetCommentDislike(idPost, idComment string, n int) {
	_, err := s.db.Exec(`INSERT INTO comment_like_dislike (dislike, post_id, user_id, comment_id) VALUES (?, ?, ?, ?)`, 1, idPost, n, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) UpdateCommentLike(idPost, idComment string, n int) {
	_, err := s.db.Exec(`UPDATE comment_like_dislike SET like=1, dislike=NULL WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) UpdateCommentDislike(idPost, idComment string, n int) {
	_, err := s.db.Exec(`UPDATE comment_like_dislike SET like=NULL, dislike=1 WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *SqlLikeStorage) CommentDislike(n int, idPost, idComment string) int {
	a := 0
	stmt := s.db.QueryRow(`SELECT "dislike" FROM "comment_like_dislike" WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	stmt.Scan(&a)
	return a
}
func (s *SqlLikeStorage) CommentLike(n int, idPost, idComment string) int {
	a := 0
	stmt := s.db.QueryRow(`SELECT "like" FROM "comment_like_dislike" WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	stmt.Scan(&a)
	return a
}
func (s *SqlLikeStorage) DeleteCommentLikeNDis(idPost, idComment string, n int) {
	_, err := s.db.Exec(`DELETE FROM comment_like_dislike WHERE user_id=? AND comment_id=? AND post_id=?`, n, idComment, idPost)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (s *SqlLikeStorage) UpdateCommentLikeCount(idPost, idComment string) {
	a := 0
	Comrow := s.db.QueryRow(`SELECT COUNT(like) FROM comment_like_dislike WHERE post_id=? AND comment_id=?`, idPost, idComment)
	e := Comrow.Scan(&a)
	if e != nil {
		return
	}

	_, err := s.db.Exec(`UPDATE comment SET likes=? WHERE id=?`, a, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *SqlLikeStorage) UpdateCommentDislikeCount(idPost, idComment string) {
	a := 0
	Comrow := s.db.QueryRow(`SELECT COUNT(dislike) FROM comment_like_dislike WHERE post_id=? AND comment_id=?`, idPost, idComment)
	e := Comrow.Scan(&a)
	if e != nil {
		return
	}

	_, err := s.db.Exec(`UPDATE comment SET dislikes=? WHERE id=?`, a, idComment)
	if err != nil {
		fmt.Println(err)
		return
	}
}
