package models

import "juro-go/database"

type Recipe struct {
	database.DefaultModel
	title string `json:"title"`
}
