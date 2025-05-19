package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"os"
	"time"
)

type JsonStorage struct {
	Tasks    []todo.Task
	FilePath string
}

func (j *JsonStorage) GetTask() error {
	data, err := os.ReadFile(j.FilePath)
	if os.IsNotExist(err) {
		err = os.WriteFile(j.FilePath, []byte("[]"), 0644)
		if err != nil {
			return err
		}
		j.Tasks = make([]todo.Task, 0)
	} else {
		err = json.Unmarshal(data, &j.Tasks)
		if err != nil {
			return err
		}
	}
	return nil
}
func (j *JsonStorage) SetTask() error {
	if dataf, err := json.Marshal(j.Tasks); err != nil {
		return err
	} else {
		if err = os.WriteFile(j.FilePath, dataf, 0644); err != nil {
			return err
		} else {
			j.Tasks = make([]todo.Task, 0)
			return nil
		}
	}
}
func (j *JsonStorage) TaskAdd(task todo.Task) error {
	for _, v := range j.Tasks {
		if v.ID == task.ID {
			return errors.New("ayni id ile task eklenemez")
		}
	}
	j.Tasks = append(j.Tasks, task)
	return nil
}
func (j *JsonStorage) TaskRemove(no int) error {
	var condition bool = false
	for i := 1; i <= len(j.Tasks); i++ {
		if i == no {
			condition = true
			break
		}
	}
	if !condition {
		return errors.New("task is not found")
	}
	j.Tasks = append(j.Tasks[:no-1], j.Tasks[no:]...)
	return nil
}
func (j *JsonStorage) TaskChange(task todo.Task, no int) error {
	var condition bool = false
	for i := 1; i <= len(j.Tasks); i++ {
		if i == no {
			condition = true
			break
		}
	}
	if !condition {
		return errors.New("task is not found")
	}
	j.Tasks[no-1] = task
	return nil
}
func (j *JsonStorage) TaskMarkDone(no int) error {
	var condition bool = false
	for i := 1; i <= len(j.Tasks); i++ {
		if i == no {
			condition = true
			break
		}
	}
	if !condition {
		return errors.New("task is not found")
	}
	j.Tasks[no-1].Done = true
	return nil
}
func (j *JsonStorage) TaskList(time time.Time) error {
	var tasks []todo.Task
	for _, v := range j.Tasks {
		if v.Date.Format("2006-01-02") == time.Format("2006-01-02") {
			tasks = append(tasks, v)
		}
	}
	if len(tasks) == 0 {
		return errors.New("bu tarihte task yok")
	}
	fmt.Println("Filetred Tasks")
	for i, v := range tasks {
		fmt.Println("Task ", i+1)
		fmt.Println("****************")
		fmt.Println("ID: ", v.ID)
		fmt.Println("Title: ", v.Title)
		fmt.Println("Content: ", v.Content)
		fmt.Println("Status: ", v.Done)
		fmt.Println("Date: ", v.Date.Format("2006-01-02 15:04"))
		fmt.Println("****************")
	}
	return nil
}
