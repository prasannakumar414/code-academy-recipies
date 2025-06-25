package models

type Recipie struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	PrepTime     string `json:"prep_time"`
}
