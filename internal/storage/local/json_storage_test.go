package local

import (
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"os"
	"testing"
	"time"
)

func TestGetSetTask(t *testing.T) {
	fmt.Println("Get Set Test")
	var setstorage JsonStorage = JsonStorage{
		FilePath: "task.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}

	err := setstorage.GetTask()
	if err != nil {
		t.Error("get hatasi setstorage", err)
	}
	fmt.Println("set oncesi setstorage: ", setstorage.Tasks)
	err = setstorage.SetTask()
	if err != nil {
		t.Error("Set hatasi", err)
	}
	fmt.Println("set sonrasi setstorage: ", setstorage.Tasks)

	var storage JsonStorage
	storage.FilePath = "task.json"

	err = storage.GetTask()
	if err != nil {
		t.Error("get hatasi storage", err)
	}
	storage.Tasks = make([]todo.Task, 0)
	err = storage.SetTask()
	if err != nil {
		t.Error(err)
	}

	data, err := os.ReadFile(storage.FilePath)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("bos json yazıldıktan sonra data", string(data))

	defer func() {
		if err = os.Remove("task.json"); err != nil {
			t.Error("cıkis hatasi")
		}
	}()
}

func TestJsonStorage_TaskList(t *testing.T) {
	fmt.Println("Task List Testinng")

	var storage JsonStorage = JsonStorage{
		FilePath: "test.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}
	err := storage.TaskList(time.Now())
	if err != nil {
		t.Error("listeleme hatasi: ", err)
	}
	err = storage.TaskList(time.Now().AddDate(0, 0, 3))
	if err != nil {
		t.Error("listeleme hatasi: ", err)
	}
	err = storage.TaskList(time.Now().AddDate(1, 0, 3))
	if err != nil {
		fmt.Println("listeleme hatasi: ", err)
	}
}

func TestJsonStorage_TaskMarkDone(t *testing.T) {
	fmt.Println("Task Mark Done Test")
	var storage JsonStorage = JsonStorage{
		FilePath: "test.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}

	err := storage.TaskMarkDone(0)
	if err != nil {
		t.Error("Mark Done error: ", err)
	}
	err = storage.TaskMarkDone(5)
	if err != nil {
		t.Log("Mark Done error: ", err)
	}
	fmt.Println(storage.Tasks[0].ID, storage.Tasks[0].Done)
	fmt.Println(storage.Tasks[1].ID, storage.Tasks[1].Done)
	fmt.Println(storage.Tasks[2].ID, storage.Tasks[2].Done)
}
func TestJsonStorage_TaskAdd(t *testing.T) {
	fmt.Println("Task Add Test")
	var storage JsonStorage = JsonStorage{
		FilePath: "test.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}
	var newTask []todo.Task = []todo.Task{
		{
			ID:      1,
			Title:   "Mutfak",
			Content: "Ekmek al",
			Date:    time.Now(),
			Done:    false,
		},
		{
			ID:      4,
			Title:   "Mutfak",
			Content: "Ekmek al",
			Date:    time.Now(),
			Done:    false,
		},
	}
	err := storage.TaskAdd(newTask[0])
	if err != nil {
		t.Log("Add error: ", err)
	}
	err = storage.TaskAdd(newTask[1])
	if err != nil {
		t.Error("Add error: ", err)
	}
	fmt.Println(storage)
}
func TestJsonStorage_TaskRemove(t *testing.T) {
	fmt.Println("Task Remove Test")
	var storage JsonStorage = JsonStorage{
		FilePath: "test.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}
	err := storage.TaskRemove(1)
	if err != nil {
		t.Error("Remove error: ", err)
	}
	err = storage.TaskRemove(1)
	if err != nil {
		t.Log("Remove error: ", err)
	}
	err = storage.TaskRemove(5)
	if err != nil {
		t.Log("Remove error: ", err)
	}
	fmt.Println(storage)
}
func TestJsonStorage_TaskChange(t *testing.T) {
	fmt.Println("task Change Test")
	var storage JsonStorage = JsonStorage{
		FilePath: "test.json",
		Tasks: []todo.Task{
			{
				ID:      0,
				Title:   "Ders",
				Content: "Matemetaki çalış",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      1,
				Title:   "Mutfak",
				Content: "Ekmek al",
				Date:    time.Now(),
				Done:    false,
			},
			{
				ID:      2,
				Title:   "Buluşma",
				Content: "çiçek almayı unutma",
				Date:    time.Now().AddDate(0, 0, 3),
				Done:    false,
			},
		},
	}
	var newTask todo.Task = storage.Tasks[0]
	newTask.Done = true
	newTask.Content = "Fizik çalış"
	err := storage.TaskChange(newTask, 0)
	if err != nil {
		t.Error("Change error: ", err)
	}
	err = storage.TaskChange(newTask, 5)
	if err != nil {
		t.Log("Change error: ", err)
	}
	err = storage.TaskChange(todo.Task{}, 1)
	if err != nil {
		t.Error("Change error: ", err)
	}
	fmt.Println(storage)
}
