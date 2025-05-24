package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoDB struct {
	DB *mongo.Database
}

func (m MongoDB) CreateTask(task Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	collection := m.DB.Collection("tasks")
	task.Done = false
	task.CreatedAt = time.Now()

	res, err := collection.InsertOne(ctx, task)
	if err != nil {
		return err
	} else {
		fmt.Println("creation is successful", res.InsertedID)
		return nil
	}
}

func (m MongoDB) GetTasksByUserId(id primitive.ObjectID) ([]Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	collection := m.DB.Collection("tasks")
	cur, err := collection.Find(ctx, bson.M{"user_id": id})
	var data []Task
	defer func() {
		cancel()
		closeErr := cur.Close(context.Background())
		if closeErr != nil {
			fmt.Println(closeErr.Error())
		}
	}()
	if err != nil {
		return nil, err
	} else {
		for cur.Next(ctx) {
			var temp Task
			if err = cur.Decode(&temp); err != nil {
				return nil, err
			}
			data = append(data, temp)
		}
		fmt.Println("getting tasks by user id is successful")
		return data, err
	}
}

func (m MongoDB) GetTaskByID(id primitive.ObjectID) (Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	collection := m.DB.Collection("tasks")
	var data Task
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&data)
	if err != nil {
		return Task{}, err
	}
	fmt.Println("Getting task by user id is successful")
	return data, nil
}

func (m MongoDB) UpdateTask(id primitive.ObjectID, task Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	task.Id = id
	collection := m.DB.Collection("tasks")
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": id}, task)
	if err != nil {
		return err
	}
	fmt.Println("updating task is successful")
	return nil
}

func (m MongoDB) DeleteTask(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	collection := m.DB.Collection("tasks")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Println("deleting task is successful")
	return nil
}
