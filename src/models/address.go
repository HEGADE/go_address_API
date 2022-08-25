package models

// State..
type Mst_state struct {
	ID   int64  `json:"ID"`
	Name string `json:"Name"`
}

// City...
//type City struct {
//	ID   int    `json:"ID"`
// 	Name string `json:"Name"`
// 	State     State  `json:"State" gorm:"foreignKey:ID"`
// }
