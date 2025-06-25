package models

type Recipie struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	Instructions string `json:"instructions" bson:"instructions"`
	PrepTime     string `json:"prep_time" bson:"prep_time"`
}