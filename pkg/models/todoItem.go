package models

import (
	"github.com/google/uuid"
)

type TaskStatus int

type TodoItems struct {
	TodoItems []TodoItem `json:"todoitems"`
}

type TodoItem struct {
	Id     string     `json:"id"`
	Group  string     `json:"group"`
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"`
}

const (
	Pending TaskStatus = iota
	Done
)

func NewTodoItem(group string, task string) *TodoItem {
	var item TodoItem
	item.Id = uuid.NewString()
	item.Group = group
	item.Task = task
	item.Status = Pending

	return &item
}
