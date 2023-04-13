package models

import "gorm.io/gorm"

type Kucing struct {
	gorm.Model
	Nama string `json:"nama"`
	Type string `json:"type"`
}
