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

func (b *Todo) updateTitle(title string) {
	(*b).title = title
}

// updateCompleted
func (b *Todo) updateCompleted() {
	b.completed = !b.completed

	if b.completed {
		now := time.Now()
		b.completedAt = &now
	} else {
		b.completedAt = nil
	}
}
