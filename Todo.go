package main

import "time"

// Todo serves as a todo-model
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a sliec of Todo
type Todos []Todo
