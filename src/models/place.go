package models

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name    string
	StateID int
	State   State
}

type State struct {
	ID   int
	Name string
}
