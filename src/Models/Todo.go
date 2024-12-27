// Created by @soulaimaneyh
// 2024-11-10

package Models

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo
