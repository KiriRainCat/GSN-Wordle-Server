package model

import "time"

type Word struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;autoIncrement"`
	Active     bool   `json:"active" gorm:"default:false"`
	Subject    string `json:"subject,omitempty"`
	Value      string `json:"word,omitempty" gorm:"unique"`
	Definition string `json:"definition,omitempty"`
	Length     int    `json:"length,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
