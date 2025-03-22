package todo

import (
	"time"
)

type Store interface {
	GetTask() error
	SetTask() error
	TaskAdd(t Task) error
	TaskRemove(i int) error
	TaskChange(t Task, i int) error
	TaskMarkDone(i int) error
	TaskList(t time.Time) error
}
