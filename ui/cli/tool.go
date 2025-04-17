package cli

import (
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"time"
)

func ShowData(db *storage.JsonStorage) {
	fmt.Println("==== ToDo APP ====")
	for i, v := range db.Tasks {
		fmt.Println("Task ", i)
		fmt.Println("----------------")
		fmt.Println("ID: ", v.ID)
		fmt.Println("Title: ", v.Title)
		fmt.Println("Content: ", v.Content)
		fmt.Println("Status: ", v.Done)
		fmt.Println("Date: ", v.Date.Format("2006-01-02"))
		fmt.Println("----------------")
	}
}
func ShowPanel() {
	fmt.Println("Please choose one of the options below:")
	fmt.Println("0 - Add Task")
	fmt.Println("1 - Remove Task")
	fmt.Println("2 - Change a Task")
	fmt.Println("3 - Filter The Tasks")
	fmt.Println("4 - Mark Done")
}
func AddTask(db *storage.JsonStorage) {
	var task todo.Task = todo.Task{}
	var ListID []int
	for _, v := range db.Tasks {
		ListID = append(ListID, v.ID)
	}
	task.Date = time.Now()
	task.ID = len(ListID)
	task.Done = false
	var title, content string
	fmt.Print("Title giriniz: ")
	if _, err := fmt.Scanf("%s", &title); err != nil {
		fmt.Println("hata: ", err)
	}
	fmt.Print("Content giriniz: ")
	if _, err := fmt.Scanf("%s", &content); err != nil {
		fmt.Println("hata: ", err)
	}
	task.Title = title
	task.Content = content
	err := db.TaskAdd(task)
	if err != nil {
		fmt.Println("hata: ", err)
	} else {
		fmt.Println("Task eklendi")
		fmt.Println(db.Tasks)
	}
	// title content alma kısmında metin tam olarak doğru alınmıyor
}
