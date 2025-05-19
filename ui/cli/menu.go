package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage/local"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var db local.JsonStorage = local.JsonStorage{
	[]todo.Task{},
	"",
}

func init() {
	path, err := filepath.Abs(filepath.Dir(os.Args[0])) //this is for binary not main
	if err != nil {
		fmt.Println("data dosya yolu hatası", err)
	}
	path = filepath.Join(path, "data.json")
	db.FilePath = path
}
func Menu() {
	if err := db.GetTask(); err != nil {
		fmt.Println(err)
	}
	var response int
	var res string
	var err error
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	for {
		ShowData(&db)
		ShowPanel()

		for {
			if res, err = reader.ReadString('\n'); err != nil {
				fmt.Println("Try Again, ERROR: ", err)
				continue
			}
			if response, err = strconv.Atoi(strings.TrimSpace(res)); err != nil {

				fmt.Println("Response has not been converted", err)
				continue
			}
			if response >= 0 && response < 6 {
				break
			}
			fmt.Println("Try Again, ERROR: ", errors.New("response must be in 0 between 5"))
		}
		switch response {
		case 0:
			task := CreateTask()
			err := db.TaskAdd(task)
			if err != nil {
				fmt.Println("hata: ", err)
			} else {
				fmt.Println("Task eklendi")
			}
		case 1:
			fmt.Println("enter a task number for remove process: ")
			for {
				if res, err = reader.ReadString('\n'); err != nil {
					fmt.Println("error: ", err)
					continue
				} else {
					id, err := strconv.Atoi(strings.TrimSpace(res))
					if err != nil {
						fmt.Println("error: ", err)
						continue
					}
					if err = db.TaskRemove(id); err != nil {
						fmt.Println("error: ", err)
						continue
					} else {
						fmt.Println("task was deleted")
						break
					}
				}
			}
		case 2:
			fmt.Println("enter a task number to Change task")
			for {
				if res, err = reader.ReadString('\n'); err != nil {
					fmt.Println("error: ", err)
					continue
				} else {
					if id, err := strconv.Atoi(strings.TrimSpace(res)); err != nil {
						fmt.Println("error: ", err)
						continue
					} else {
						fmt.Println("new task: ")
						if err = db.TaskChange(CreateTask(), id); err != nil {
							fmt.Println("error: ", err)
							continue
						}
					}
				}
				fmt.Println("Task changed succesfully!")
				break
			}
		case 3:
			time := TakeDate()
			if err = db.TaskList(time); err != nil {
				fmt.Println("error: ", err)
			}
		case 4:
			fmt.Println("enter a task number to mark done")
			for {
				if res, err = reader.ReadString('\n'); err != nil {
					fmt.Println("error: ", err)
					continue
				} else {
					if id, err := strconv.Atoi(strings.TrimSpace(res)); err != nil {
						fmt.Println("error: ", err)
						continue
					} else {
						if err = db.TaskMarkDone(id); err != nil {
							fmt.Println("error: ", err)
							continue
						}
					}
				}
				fmt.Println("Task-Marking was done succesfully!")
				break
			}
		case 5:
			fmt.Println("Exiting..")
			if err = db.SetTask(); err != nil {
				fmt.Println("error: ", err)
				continue
			}
			return
		}
	}

}
