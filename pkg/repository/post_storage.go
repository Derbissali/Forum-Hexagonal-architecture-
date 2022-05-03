package repository

import (
	"database/sql"
	"fmt"
	"log"
	"tidy/pkg/model"
)

type SqlPostStorage struct {
	db *sql.DB
}

func NewPostStorage(db *sql.DB) *SqlPostStorage {
	return &SqlPostStorage{
		db: db,
	}
}
func (c *SqlPostStorage) GetAll() ([]model.Post, error) {
	rows, err := c.db.Query(`SELECT post.id, user.name, post.name, post.body, post.Image
	FROM post
	INNER JOIN user ON user.id=post.user_id ORDER BY post.id DESC`)
	var m model.Post
	if err != nil {
		log.Println(err)
		return m.Rows, err
	}

	for rows.Next() {
		var a model.Post
		err = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if err != nil {
			log.Println(err)
			return m.Rows, nil
		}
		a.Cat = CategoryByID(c, a.ID)
		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func (c *SqlPostStorage) GetSearch(str string) ([]model.Post, error) {
	fmt.Println(str)
	rows, err := c.db.Query(`SELECT post.id, user.name, post.name, post.body, post.Image FROM post
	 INNER JOIN user ON user.id=post.user_id  WHERE instr(post.name, ?) ORDER BY post.id DESC`, str)
	var m model.Post
	if err != nil {
		log.Println(err)
		return m.Rows, err
	}

	for rows.Next() {
		var a model.Post
		err = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if err != nil {
			log.Println(err)
			return m.Rows, nil
		}
		a.Cat = CategoryByID(c, a.ID)
		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func CategoryByID(c *SqlPostStorage, postid int) []model.Category {
	row, e := c.db.Query(`SELECT category.name  
	FROM category_post
	INNER JOIN category ON category.id = category_post.category_id
	WHERE post_id = ?`, postid)
	if e != nil {
		log.Println(e)
		return nil
	}
	var m model.Post
	defer row.Close()
	for row.Next() {
		var a model.Category
		e = row.Scan(&a.Name)
		if e != nil {
			log.Println(e)
			return nil
		}

		m.Cat = append(m.Cat, a)

	}
	return m.Cat
}

func (c *SqlPostStorage) GetCategory() ([]model.Category, error) {
	rows, e := c.db.Query(`SELECT "name" FROM "category" ORDER BY "name"`)
	if e != nil {
		return nil, e
	}
	var m model.Category
	for rows.Next() {
		cat1 := model.Category{}
		e = rows.Scan(&cat1.Name)
		if e != nil {
			return nil, e
		}
		m.Rows = append(m.Rows, cat1)

	}
	return m.Rows, nil
}
func (c *SqlPostStorage) SortedCategory(t string) ([]model.Post, error) {
	rows, err := c.db.Query((`SELECT post.id, user.name, post.name, post.body, post.Image
	FROM category 
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN user ON user.id=post.user_id
	WHERE category.name = ? ORDER BY post.id DESC`), t)
	var m model.Post
	if err != nil {
		log.Println(err)
		return m.Rows, err
	}

	for rows.Next() {
		var a model.Post
		err = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if err != nil {
			log.Println(err)
			return m.Rows, nil
		}
		a.Cat = CategoryByID(c, a.ID)
		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func (c *SqlPostStorage) LikedPosts(t int) ([]model.Post, error) {

	rows, err := c.db.Query(`SELECT DISTINCT post.id, user.name, post.name, post.body, post.Image
	FROM category 
	INNER JOIN likeNdis ON post.id = likeNdis.post_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN user ON user.id = likeNdis.user_id WHERE likeNdis.like=1 AND user.id=? ORDER BY post.id DESC`, t)
	var m model.Post
	if err != nil {
		log.Println(err)
		return m.Rows, err
	}

	for rows.Next() {
		var a model.Post
		err = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if err != nil {
			log.Println(err)
			return m.Rows, nil
		}
		a.Cat = CategoryByID(c, a.ID)
		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func (c *SqlPostStorage) CreatedPosts(t int) ([]model.Post, error) {
	fmt.Println(t, "asd")
	rows, err := c.db.Query(`SELECT DISTINCT post.id, user.name, post.name, post.body, post.Image
	FROM category 
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN user ON user.id = post.user_id WHERE user.id=? ORDER BY post.id DESC`, t)
	var m model.Post
	if err != nil {
		log.Println(err)
		return m.Rows, err
	}

	for rows.Next() {
		var a model.Post
		err = rows.Scan(&a.ID, &a.User.Name, &a.Name, &a.Body, &a.Image)
		if err != nil {
			log.Println(err)
			return m.Rows, nil
		}
		a.Cat = CategoryByID(c, a.ID)
		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
func (c *SqlPostStorage) CreatePost(m model.Post, s []string, ID int) error {
	b := 0
	stmt := c.db.QueryRow(`INSERT INTO post (name, body, user_id, image, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?) RETURNING id`, m.Name, m.Body, ID, m.Image, 0, 0)
	fmt.Println(s)
	stmt.Scan(&b)
	for _, i := range s {
		a := 0
		stmt1 := c.db.QueryRow(`SELECT "id" FROM "category" WHERE category.name=?`, i)
		stmt1.Scan(&a)
		if a == 0 {
			fmt.Println("BadRequest")
			return nil
		}
		_, err := c.db.Exec(`INSERT INTO category_post (category_id, post_id) VALUES (?, ?) `, a, b)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	return nil
}

func (c *SqlPostStorage) CountPost() (int, error) {
	a := 0
	row := c.db.QueryRow(`SELECT COUNT(DISTINCT name) FROM post`)
	err := row.Scan(&a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (c *SqlPostStorage) SinglePost(id int) ([]model.Post, error) {
	row := c.db.QueryRow((`SELECT post.id, post.name, post.body, post.likes, post.dislikes, user.Name, post.Image
	FROM category 
	INNER JOIN category_post ON category.id = category_post.category_id
	INNER JOIN post ON post.id = category_post.post_id
	INNER JOIN user ON user.id=post.user_id
	WHERE post.id = ?`), id)
	var a model.Post
	var m model.Post
	err := row.Scan(&a.ID, &a.Name, &a.Body, &a.Likes, &a.Dislikes, &a.User.Name, &a.Image)
	if err != nil {
		return a.Rows, err
	}
	a.Cat = CategoryByID(c, a.ID)

	a.Comm, err = c.GetAllComment(a.ID)
	if err != nil {
		return m.Rows, err
	}

	m.Rows = append(m.Rows, a)

	return m.Rows, nil
}
func (s *SqlPostStorage) GetAllComment(id int) ([]model.Comment, error) {
	rows, err := s.db.Query((`SELECT comment.id, body, post_id, user.name, comment.likes, comment.dislikes
	FROM comment 
	INNER JOIN user ON user.id = comment.user_id
	WHERE post_id=?`), id)
	if err != nil {
		return nil, err
	}
	var m model.Comment
	for rows.Next() {
		var a model.Comment
		err = rows.Scan(&a.ID, &a.Body, &a.Post.ID, &a.User.Name, &a.Likes, &a.Dislikes)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		m.Rows = append(m.Rows, a)

	}

	return m.Rows, nil
}
