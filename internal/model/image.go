package model

import "database/sql"

type Image struct {
		Me     string `json:"me"`
	Friend string `json:"friend"`
}

type ImageModel struct {
	DB *sql.DB
}
