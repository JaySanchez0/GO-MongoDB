package Persistence

import (
	"context"
	"fmt"
	"log"
	"main/model"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PeoplePersistence struct {
	Db *mongo.Database
}
/**
* Devuelve todas las personas en la base de datos
*/
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
/**
* Registra una persona en la base de datos
*/
func (p PeoplePersistence) AddPeople(people model.People) {
	db := p.FillDefaults()
	fmt.Println("save")
	collection := db.Collection("People")
	collection.InsertOne(context.TODO(), people.Clone())
}

/**
* obtiene la conexion a la base de datos
*/
func (p PeoplePersistence) FillDefaults() *mongo.Database {
	if p.Db == nil {
		//fmt.Println("entro create peristence")
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test:testApp123@cluster0.2xxny.mongodb.net/PeopleDB?retryWrites=true&w=majority").SetWriteConcern(writeconcern.New(writeconcern.WMajority())))
		//client, err := mongo.Connect(context.TODO(), clientOptions)
		if err == nil {
			err = client.Connect(context.TODO())
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
