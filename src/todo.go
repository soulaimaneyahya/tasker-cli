/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import (
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

// update title by index
func (todos *Todos) updateTitleByIndex(index int, title string) error {
	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].title = title
	return nil
}

// update completedAt by index
func (todos *Todos) updateCompletedAtByIndex(index int) error {
	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index]
	todo.completed = !todo.completed

	if todo.completed {
		now := time.Now()
		todo.completedAt = &now
	} else {
		todo.completedAt = nil
	}

	return nil
}

// delete todo
func (todos *Todos) deleteTodo(index int) error {
	if err := (*todos).validateIndex(index); err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	return nil
}
