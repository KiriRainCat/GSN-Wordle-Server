package model

type Word struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;autoIncrement"`
	Value      string `json:"value,omitempty" gorm:"unique"`
	Definition string `json:"definition,omitempty"`
	Length     int    `json:"length,omitempty"`
}