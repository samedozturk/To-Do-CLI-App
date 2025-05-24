package todo

import (
	"github.com/samedozturk/To-Do-CLI-App/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type JsonStore interface {
	GetTask() error
	SetTask() error
	TaskAdd(t Task) error
	TaskRemove(i int) error
	TaskChange(t Task, i int) error
	TaskMarkDone(i int) error
	TaskList(t time.Time) error
}
type MongoStore interface {
	CreateTask(task mongodb.Task) error

	GetTasksByUserId(id primitive.ObjectID) ([]mongodb.Task, error)
	GetTaskByID(id primitive.ObjectID) (mongodb.Task, error)

	UpdateTask(id primitive.ObjectID, task mongodb.Task) error

	DeleteTask(id primitive.ObjectID) error
}
