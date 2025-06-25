package models

type Recipie struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	Instructions string `json:"instructions" bson:"instructions"`
	PrepTime     string `json:"prep_time" bson:"prep_time"`
}

type RecipiePagination struct {
	Recipies []Recipie `json:"recipies"`
	Meta RecipiePaginationMeta `json:"meta" bson:"meta"`
}

type RecipiePaginationMeta struct {
	Total int `json:"total" bson:"total"`
	Skipped int `json:"skipped" bson:"skipped"`
	Limit int `json:"limit"`
}

type RecipiePaginationRequest struct {
	Skip int `json:"skip"`
	Limit int `json:"limit"`
	Query string `json:"query"`
}