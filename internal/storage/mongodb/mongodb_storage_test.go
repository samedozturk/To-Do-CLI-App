package mongodb

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func prepareTestDb(t *testing.T) MongoDB {
	t.Helper()

	_ = godotenv.Load("C:\\Users\\USER\\Desktop\\Go\\Projects\\To-Do-CLI-App\\.env")
	db := MongoDB{DB: SetupDB("exampleForTest")}

	t.Cleanup(func() {
		err := db.DB.Drop(context.Background())
		t.Log("drop db error: ", err)
	})
	return db
}
func generateTask() Task {
	task := Task{
		Title:      "ders",
		Content:    "matematik çalış",
		UserId:     primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		Done:       false,
		ExpireDate: time.Now().AddDate(0, 0, 2),
	}
	return task
}

func TestMongoDB_CreateTask(t *testing.T) {
	db := prepareTestDb(t)
	task := generateTask()

	if err := db.CreateTask(task); err != nil {
		t.Error("create error", err.Error())
	}
}
func TestMongoDB_GetTasksByUserId(t *testing.T) {
	db := prepareTestDb(t)
	userId := primitive.NewObjectID()
	task0 := generateTask()
	task1 := generateTask()
	task0.UserId = userId
	task1.UserId = userId
	err := db.CreateTask(task0)
	if err != nil {
		t.Error("creation error", err.Error())
	}
	err = db.CreateTask(task1)
	if err != nil {
		t.Error("creation error:", err.Error())
	}

	returnedTask, err := db.GetTasksByUserId(userId)
	if err != nil {
		t.Error("get task by user id process' error:", err.Error())
	}
	for _, v := range returnedTask {
		fmt.Println("user id:", v.UserId)
		fmt.Println("task id:", v.Id)
	}
}
func TestMongoDB_GetTaskByID(t *testing.T) {
	db := prepareTestDb(t)
	var id primitive.ObjectID = primitive.NewObjectID()
	task := generateTask()
	task.Id = id

	if err := db.CreateTask(task); err != nil {
		t.Error("create error", err.Error())
	}
	tsk, err := db.GetTaskByID(task.Id)
	if err != nil {
		t.Error("get task by task id error:", err.Error())
	}
	fmt.Println(tsk)
}

func TestMongoDB_UpdateTask(t *testing.T) {
	db := prepareTestDb(t)
	task1 := generateTask()
	taskId1 := primitive.NewObjectID()
	taskId2 := primitive.NewObjectID()
	task1.Id = taskId1
	err := db.CreateTask(task1)
	if err != nil {
		t.Error("create error:", err.Error())
	}
	fmt.Println(db.GetTaskByID(task1.Id))

	task2 := generateTask()
	task2.Id = taskId2
	task2.Title = "different task"

	err = db.UpdateTask(task1.Id, task2)
	if err != nil {
		t.Error("update error:", err.Error())
	}
	fmt.Println(db.GetTaskByID(task1.Id))
}

func TestMongoDB_DeleteTask(t *testing.T) {
	db := prepareTestDb(t)
	task := generateTask()
	task.Id = primitive.NewObjectID()
	err := db.CreateTask(task)
	if err != nil {
		t.Error("create error:", err.Error())
	}

	err = db.DeleteTask(task.Id)
	if err != nil {
		t.Error("delete error:", err.Error())
	}
}
