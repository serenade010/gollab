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

func (m *FriendListlModel) GetList(user string) ([]string, error) {
	stmt := "SELECT friend FROM gollab_friendlist WHERE me=$1"
	rows, err := m.DB.Query(stmt, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var friends []string
	for rows.Next() {
		var friend string
		err := rows.Scan(&friend)
		if err != nil {
			return nil, err
		}

		friends = append(friends, friend)

	}

	return friends, nil
}
