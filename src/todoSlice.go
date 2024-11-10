/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aquasecurity/table"
)

// validate Todo slice index
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// appendTodo adds a new Todo to the Todo slice.
func (todos *Todos) appendTodo(x1todo *Todo) {
	*todos = append(*todos, *x1todo)
}

// Receiver function to format Todos as a Markdown table with aligned fields
func (todos *Todos) format() string {
	t := table.New(os.Stdout)
	t.SetHeaders("No", "Title", "Status", "Created At", "Completed At")
	t.SetHeaderAlignment(table.AlignCenter)
	t.SetAlignment(table.AlignLeft)

	if len(*todos) < 1 {
		return "No todo found"
	}

	for i, todo := range *todos {
		status := "Incomplete"
		completedAt := "null"

		if todo.completed {
			status = "Completed"
			if todo.completedAt != nil {
				completedAt = todo.completedAt.Format("2006-01-02 15:04:05")
			}
		}

		t.AddRow(
			fmt.Sprintf("%d", i+1),
			todo.title,
			status,
			todo.createdAt.Format("2006-01-02 15:04:05"),
			completedAt,
		)
	}

	t.Render()
	return ""
}
