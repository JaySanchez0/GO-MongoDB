package Persistence

import (
	"context"
	"fmt"
	"log"
	"main/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PeoplePersistence struct {
	Db *mongo.Database
}

func (p PeoplePersistence) GetPeoples() []model.People {
	db := p.FillDefaults()
	collection := db.Collection("People")
	var peoples []model.People
	cursor, er := collection.Find(context.TODO(), bson.D{})
	if er == nil {
		for cursor.Next(context.TODO()) {
			//Create a value into which the single document can be decoded
			var elem model.People
			err := cursor.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			peoples = append(peoples, elem)

		}
	}
	return peoples
}

func (p PeoplePersistence) AddPeople(people *model.People) {
	db := p.FillDefaults()
	collection := db.Collection("People")
	collection.InsertOne(context.TODO(), people)
}

func (p PeoplePersistence) FillDefaults() *mongo.Database {
	if p.Db == nil {
		//fmt.Println("entro create peristence")
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test:testApp123@cluster0.2xxny.mongodb.net/PeopleDB?retryWrites=true&w=majority"))
		//client, err := mongo.Connect(context.TODO(), clientOptions)
		if err == nil {
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			}
			db := client.Database("PeopleDB")
			p.Db = db
			fmt.Println("Ok")
			fmt.Println(db)
			return db
		}
		fmt.Println("Error")
	}
	return nil
}
