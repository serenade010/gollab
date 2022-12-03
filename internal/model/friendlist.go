package model

import (
	"database/sql"
)

type FriendList struct {
	Me     string `json:"me"`
	Friend string `json:"friend"`
}

type FriendListlModel struct {
	DB *sql.DB
}

func (m *FriendListlModel) Insert(me string, friend string) error {
	stmt := "INSERT INTO gollab_friendlist (me,friend) values ($1,$2)"
	_, err := m.DB.Exec(stmt, me, friend)
	if err != nil {
		return err
	}
	return nil
}
