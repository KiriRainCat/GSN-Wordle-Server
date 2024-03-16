package model

type Word struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;autoIncrement"`
	Subject    string `json:"subject,omitempty"`
	Value      string `json:"word,omitempty" gorm:"unique"`
	Definition string `json:"definition,omitempty"`
	Length     int    `json:"length,omitempty"`
}
