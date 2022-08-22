package models

// State..
type State struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

// City...
//type City struct {
//	ID   int    `json:"ID"`
// 	Name string `json:"Name"`
// 	State     State  `json:"State" gorm:"foreignKey:ID"`
// }
