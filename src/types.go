/**
* Created by @soulaimaneyh on 2024-11-10
 */

package main

import "time"

// Todo struct type
type Todo struct {
	Title       string
	Completed   bool
	createdAt   time.Time
	CompletedAt *time.Time
}
