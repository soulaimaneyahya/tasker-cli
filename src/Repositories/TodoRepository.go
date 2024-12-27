package Repositories

import (
	"errors"
	"tasker-cli/src/Models"
	"time"
)

type TodoRepository struct {
	Todos Models.Todos
}

// return Todos
func (repo *TodoRepository) GetTodos() []Models.Todo {
	return repo.Todos
}

// ValidateIndex checks if the index is valid for the Todos slice.
func (repo *TodoRepository) ValidateIndex(index int) error {
	if index < 0 || index >= len(repo.Todos) {
		return errors.New("invalid index")
	}
	return nil
}

// Append adds a new Todo to the repository.
func (repo *TodoRepository) Append(todo *Models.Todo) {
	repo.Todos = append(repo.Todos, *todo)
}

// UpdateTitleByIndex updates the title of a Todo at the given index.
func (repo *TodoRepository) UpdateTitleByIndex(index int, title string) error {
	if err := repo.ValidateIndex(index); err != nil {
		return err
	}
	repo.Todos[index].Title = title
	return nil
}

// UpdateCompletedAtByIndex toggles the Completed status and sets CompletedAt timestamp.
func (repo *TodoRepository) UpdateCompletedAtByIndex(index int) error {
	if err := repo.ValidateIndex(index); err != nil {
		return err
	}

	todo := &repo.Todos[index]
	todo.Completed = !todo.Completed

	if todo.Completed {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	return nil
}

// Delete removes the Todo at the specified index.
func (repo *TodoRepository) Delete(index int) error {
	if err := repo.ValidateIndex(index); err != nil {
		return err
	}

	repo.Todos = append(repo.Todos[:index], repo.Todos[index+1:]...)
	return nil
}
