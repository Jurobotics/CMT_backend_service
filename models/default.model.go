package models

import "time"

// This package contians the default structure for every table
type Default struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
}
