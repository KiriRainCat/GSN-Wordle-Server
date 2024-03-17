package model

import "time"

type Commit struct {
	Id         int    `json:"id,omitempty" gorm:"primary_key;autoIncrement"`
	WordId     int    `json:"wid,omitempty"`
	Subject    string `json:"subject,omitempty"`
	Word       string `json:"word,omitempty" gorm:"unique"`
	Definition string `json:"definition,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}
