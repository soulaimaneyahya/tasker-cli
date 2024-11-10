/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import "time"

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
