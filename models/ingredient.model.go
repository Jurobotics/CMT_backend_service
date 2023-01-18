package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name    string   `json:"name"`
	Amount  uint     `json:"amount"` // Amount of the doser in the Si unit
	ServoId []Servos `json:"servoId"`
}
