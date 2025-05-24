package main

import (
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/config"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config.LoadEnv()
	var database *mongo.Database = mongodb.SetupDB("exampleDB")
	db := mongodb.MongoDB{DB: database}
	id, _ := primitive.ObjectIDFromHex("000000000000000000000000")
	data, err := db.GetTasksByUserId(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data[1])
}
