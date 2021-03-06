package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	ID 		primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string `json:"name" form:"name" query:"name" bson:"name"`
	City    string `json:"city" form:"city" query:"city" bson:"city"`
	Age     int `json:"age" form:"age" query:"age" bson:"age"`
	Address string `json:"address" form:"address" query:"address" bson:"address"`
}
/**
* Devuelve una copia del objeto persona
*/
func (p People) Clone() People{
	return People{
		ID: p.ID,
		Name: p.Name,
		City: p.City,
		Age:p.Age,
		Address:p.Address,
	}
}
