package cli

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"os"
	"strings"
	"time"
)

func ShowData(db *storage.JsonStorage) {
	fmt.Println("==== ToDo APP ====")
	for i, v := range db.Tasks {
		fmt.Println("Task ", i+1)
		fmt.Println("----------------")
		fmt.Println("ID: ", v.ID)
		fmt.Println("Title: ", v.Title)
		fmt.Println("Content: ", v.Content)
		fmt.Println("Status: ", v.Done)
		fmt.Println("Date: ", v.Date.Format("2006-01-02 15:04"))
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
	fmt.Println("5 - Exit")
}
func CreateTask(db *storage.JsonStorage) todo.Task {
	gen := uuid.New()
	var task todo.Task = todo.Task{}
	task.Date = time.Now()
	task.ID = gen.ID()
	task.Done = false
	var title, content string
	fmt.Print("Title giriniz: ")
	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("hata: ", err)
	}
	fmt.Print("Content giriniz: ")
	content, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("hata: ", err)
	}

	task.Title = strings.TrimSpace(title)
	task.Content = strings.TrimSpace(content)
	return task
}
func TakeDate() time.Time {
	fmt.Println("Please enter a date (YY-MM-DD)(Please enter like this way)")
	reader := bufio.NewReader(os.Stdin)
	respond, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error: ", err)
	}
	respond = strings.TrimSpace(respond)
	if t, err := time.Parse("2006-01-02", respond); err != nil {
		fmt.Println("error: ", err)
		return time.Time{}
	} else {
		return t
	}
}

// id çakışması yaşıyoruz okeyy
// bu sorunu çöz ve unit testyaz menu.go için
