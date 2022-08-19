package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type City struct {
	gorm.Model
	Name    string `json:"name"`
	StateID int    `json:"StateID"`
	State   State  `json:"State"`
}

type State struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
