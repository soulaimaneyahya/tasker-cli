// Created by @soulaimaneyh
// 2024-09-14
// https://soulaimaneyahya.com

package Services

import (
	"time"

	"tasker-cli/src/Models"
)

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) Create(title string) Models.Todo {
	return Models.Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
}
