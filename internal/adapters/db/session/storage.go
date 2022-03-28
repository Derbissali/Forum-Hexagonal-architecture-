package session

import (
	"database/sql"
	"fmt"
	"tezt/hexagonal/internal/domain/session"
	"tezt/hexagonal/internal/model"
)

type sessionStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) session.SessionStorage {
	return &sessionStorage{
		db: db,
	}
}

func (h *sessionStorage) Create(uuid, id string) error {
	fmt.Println("asd", uuid, id)
	_, err := h.db.Exec(`INSERT INTO session (uuid, user_id) VALUES (?, ?)`, uuid, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h *sessionStorage) Check() error {
	return nil
}
func (h *sessionStorage) Delete(n string) error {
	_, err := h.db.Exec(`DELETE FROM session where user_id = ?`, n)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h *sessionStorage) ReadByUUID(uuid string) (model.User, error) {
	row, err := h.db.Query(`SELECT user.id, user.name
	FROM user
	INNER JOIN session ON session.user_id=user.id
	WHERE uuid = ?`, uuid)
	var m model.User
	if err != nil {
		fmt.Println(err)
		return m, nil
	}

	//	fmt.Println(uuid)
	n := model.User{}
	for row.Next() {
		e := row.Scan(&n.ID, &n.Name)
		if e != nil {
			fmt.Println(e)
			return m, e
		}

	}
	if n.ID != 0 {
		m.Session = true
	} else {
		m.Session = false

	}

	m.ID = n.ID
	m.Name = n.Name
	//	fmt.Println(m.ID, "from")
	return m, nil
}
