package models

type Challenge struct {
	ID int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Flag string `json:"flag" bson:"flag"`
}