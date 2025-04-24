package storage

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
func (j *JsonStorage) TaskRemove(id int) error {
	var index int = -1
	for i, v := range j.Tasks {
		if v.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("task bulunamadi")
	}
	j.Tasks = append(j.Tasks[:index], j.Tasks[index+1:]...)
	return nil
}
func (j *JsonStorage) TaskChange(task todo.Task, id int) error {
	var index int = -1
	for i := range j.Tasks {
		if i == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("task bulunamadi")
	}
	j.Tasks[index] = task
	return nil
}
func (j *JsonStorage) TaskMarkDone(id int) error {
	var index int = -1
	for i, v := range j.Tasks {
		if v.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("task bulunamadı")
	}
	j.Tasks[index].Done = true
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
	for _, v := range tasks {
		fmt.Println(v)
	}
	return nil
}
