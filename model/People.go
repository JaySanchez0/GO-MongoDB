package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type People struct {
	Id 		primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name    string `json:"name" form:"name" query:"name" bson:"name"`
	City    string `json:"city" form:"city" query:"city" bson:"city"`
	Age     int `json:"age" form:"age" query:"age" bson:"age"`
	Address string `json:"address" form:"address" query:"address" bson:"address"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
    UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
