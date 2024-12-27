// Created by @soulaimaneyh
// 2024-11-10

// Utility for table formatting

package Utils

import (
	"fmt"
	"os"

	"tasker-cli/src/Models"

	"github.com/aquasecurity/table"
)

type TableFormatter struct {
	Todos Models.Todos
}

func FormatTodos(tableFormatter *TableFormatter) string {
	t := table.New(os.Stdout)
	t.SetHeaders("No", "Title", "Status", "Created At", "Completed At")
	t.SetHeaderAlignment(table.AlignCenter)
	t.SetAlignment(table.AlignLeft)

	if len(tableFormatter.Todos) < 1 {
		return "No todo found"
	}

	for i, todo := range tableFormatter.Todos {
		status := "☐"
		completedAt := "null"

		if todo.Completed {
			status = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}

		t.AddRow(
			fmt.Sprintf("%d", i+1),
			todo.Title,
			status,
			todo.CreatedAt.Format("2006-01-02 15:04:05"),
			completedAt,
		)
	}

	t.Render()
	return ""
}
