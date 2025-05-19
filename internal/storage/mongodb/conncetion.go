package mongodb

import (
	"context"
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectDB(dbName string, ch chan *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.GetEnv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("MongoDB connection is not successful")
		log.Fatal(err)
	}
	db := client.Database(dbName)
	ch <- db
	close(ch)
}
func SetupDB(dbName string) *mongo.Database {
	ch := make(chan *mongo.Database, 1)
	go connectDB(dbName, ch)
	fmt.Println("MongoDB connection is successful")
	return <-ch
}
