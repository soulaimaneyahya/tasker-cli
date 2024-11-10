/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import (
	"fmt"
	"time"
)

// creates a new Todo
func newTodo(title string) Todo {
	return Todo{
		title:       title,
		completed:   false,
		createdAt:   time.Now(),
		completedAt: nil,
	}
}

// update todo title
func (x1todo *Todo) updateTitle(title string) {
	// x1todo.title = title
	(*x1todo).title = title
}

// update todo completed
func (x1todo *Todo) updateCompleted() {
	(*x1todo).completed = !(*x1todo).completed

	if (*x1todo).completed {
		now := time.Now()
		(*x1todo).completedAt = &now
	} else {
		(*x1todo).completedAt = nil
	}
}

// appendTodo adds a new Todo to the Todos slice.
func (todos *Todos) appendTodo(x1todo *Todo) {
	*todos = append(*todos, *x1todo)
}

// delete todo title
func (todos *Todos) deleteTitle(index int) {
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
}

// Receiver functions
// format method for Todos
func (todos *Todos) format() string {
	fs := "Todo List:\n"

	if len(*todos) < 1 {
		fs += "No todo found"

		return fs
	}

	for i, todo := range *todos {
		status := "Incomplete"

		if todo.completed {
			status = "Completed"
		}

		entry := fmt.Sprintf("%d. %s (Status: %s, Created At: %s", i+1, todo.title, status, todo.createdAt.Format("2006-01-02 15:04:05"))

		if todo.completed && todo.completedAt != nil {
			entry += fmt.Sprintf(", Completed At: %s", todo.completedAt.Format("2006-01-02 15:04:05"))
		}

		entry += ")\n"
		fs += entry
	}

	return fs
}
