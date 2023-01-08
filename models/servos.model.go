package models

import "gorm.io/gorm"

type Servos struct {
	gorm.Model
	Arduino int `json:"arduino"`
}
