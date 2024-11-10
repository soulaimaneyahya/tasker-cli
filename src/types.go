/**
* Created by @soulaimaneyh on 2024-11-10
 */

package main

import "time"

// Todo struct type
type Todo struct {
	title       string
	completed   bool
	createdAt   time.Time
	completedAt *time.Time
}

type Todos []Todo
