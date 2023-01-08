package models

import "juro-go/database"

type Recipe struct {
	database.DefaultModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Preparation string `json:"preparation"`
	Image       string `json:"image"`
}
