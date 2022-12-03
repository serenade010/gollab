package model

import (
	"database/sql"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}



type UserlModel struct {
	DB *sql.DB
}

func (m *UserlModel) Insert(name string) error {
	stmt := "INSERT INTO gollab_user (account) values ($1)"
	_, err := m.DB.Exec(stmt, name)
	if err != nil {
		return err
	}
	return nil
}
