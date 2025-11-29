package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int       `json:"task_id"`
	Name        string    `json:"name"`
	IsComplete  bool      `json:"is_complete"`
	Description string    `json:"description,omitempty"` // Description can be empty and omitted from the JSON
	CreatedAt   time.Time `json:"created_at"`
}

func createTask(id int, name string, description string) *Task {
	return &Task{
		Id:          id,
		Name:        name,
		Description: description,
		IsComplete:  false,
		CreatedAt:   time.Now(),
	}
}

func toString(t Task) string {
	return fmt.Sprintf("{\n"+
		"      id: %d\n"+
		"      name: %s\n"+
		"      description: %s\n"+
		"      is_complete: %t\n"+
		"      created_at: %s\n}",
		t.Id, t.Name, t.Description, t.IsComplete, t.CreatedAt)
}

func (task *Task) Complete() {
	task.IsComplete = true
}
