package todo

import (
	"time"
)

type Task struct {
	ID      uint32    `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	Done    bool      `json:"done"`
}
