package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
	"os"
	"strconv"
	"strings"
)

func Menu() {
	var db storage.JsonStorage = storage.JsonStorage{
		[]todo.Task{},
		"data/db1",
	}
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
			fmt.Println("Try Again, ERROR: ", errors.New("response must be in 0 between 4"))
		}
		switch response {
		case 0:
			task := CreateTask(&db)
			err := db.TaskAdd(task)
			if err != nil {
				fmt.Println("hata: ", err)
			} else {
				fmt.Println("Task eklendi")
			}
		case 1:
			fmt.Println("enter a id number for remove process: ")
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
						fmt.Println("task has been deleted")
						break
					}
				}
			}
		case 2:
			fmt.Println("enter a id number to Change task")
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
						if err = db.TaskChange(CreateTask(&db), id); err != nil {
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
			fmt.Println("enter a id number to mark done")
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
				fmt.Println("Task marked done succesfully!")
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
