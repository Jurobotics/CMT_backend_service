package models

import "juro-go/database"

type Servos struct {
	database.DefaultModel
	Arduino int `json:"arduino"`
}
