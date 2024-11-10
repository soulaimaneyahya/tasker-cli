/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import "fmt"

func main() {
	x1todo := newTodo("PHP")
	x1todo.updateTitle("x1PHP")
	x1todo.updateCompleted()

	x2todo := newTodo("Laravel")

	x3todo := newTodo("Symfony")

	x4todo := newTodo("Go")

	var todos Todos
	todos.appendTodo(&x1todo)
	todos.appendTodo(&x2todo)
	todos.appendTodo(&x3todo)
	todos.appendTodo(&x4todo)

	todos.deleteTitle(1)
	fmt.Println(todos.format())
}
