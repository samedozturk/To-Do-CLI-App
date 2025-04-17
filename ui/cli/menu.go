package cli

import (
	"errors"
	"fmt"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage"
	"github.com/samedozturk/To-Do-CLI-App/internal/todo"
)

func Menu() {
	var db storage.JsonStorage = storage.JsonStorage{
		[]todo.Task{},
		"github.com/samedozturk/To-Do-CLI-App/data/db1",
	}
	if err := db.GetTask(); err != nil {
		fmt.Println(err)
	}
	ShowData(&db)
	ShowPanel()
	var response int8
	for {
		if _, err := fmt.Scanf("%d", &response); err != nil {
			fmt.Println("Try Again, ERROR: ", err)
			continue
		}
		if response >= 0 && response < 5 {
			break
		}
		fmt.Println("Try Again, ERROR: ", errors.New("response must be in 0 between 4"))
	}
	switch response {
	case 0:
		AddTask(&db)
		ShowData(&db)
		ShowPanel()
	case 1:
	case 2:
	case 3:
	case 4:
		// metin girdisi olunca tek bir hata fırlatmıyor
	}
}
