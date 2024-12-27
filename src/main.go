package main

import (
	"tasker-cli/src/Models"
	"tasker-cli/src/Repositories"
	"tasker-cli/src/Utils"
)

func main() {
	repo := &Repositories.TodoRepository{}

	// Create Todos
	x1todo := &Models.Todo{Title: "PHP"}
	x2todo := &Models.Todo{Title: "Laravel"}
	x3todo := &Models.Todo{Title: "Symfony"}
	x4todo := &Models.Todo{Title: "Go"}

	repo.Append(x1todo)
	repo.Append(x2todo)
	repo.Append(x3todo)
	repo.Append(x4todo)

	repo.UpdateTitleByIndex(0, "xPHP")

	repo.UpdateCompletedAtByIndex(0)

	repo.Delete(1)

	tableFormatter := &Utils.TableFormatter{Todos: repo.GetTodos()}
	Utils.FormatTodos(tableFormatter)
}
