package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	collection := cli.Database("test").Collection("collection")

	// insert
	insertRes, err := collection.InsertOne(ctx, bson.D{{"name", "zx"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("InsertID=%s\n", insertRes.InsertedID)

	// read
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var m struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}
	for cur.Next(ctx) {
		if err = cur.Decode(&m); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", m)
	}

	// update
	updateRes, err := collection.UpdateOne(ctx,
		bson.D{{"_id", insertRes.InsertedID}},
		bson.D{
			{"$set", bson.D{{"age", 30}}},
			{"$set", bson.D{{"city", "Hangzhou"}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("modified count=%d\n", updateRes.ModifiedCount)
	findRes := collection.FindOne(ctx, bson.D{{"name", "zx"}})
	err = findRes.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated doc:%+v\n", m)

	// delete
	deleteRes, err := collection.DeleteOne(ctx, bson.D{{"name", "zx"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete count=%d\n", deleteRes.DeletedCount)
}
