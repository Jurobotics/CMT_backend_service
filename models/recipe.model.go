package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Preparation string `json:"preparation"`
	Image       string `json:"image"`
}
