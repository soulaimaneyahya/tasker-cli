/**
* Created by @soulaimaneyh on 2024-09-14
 */

package main

import "fmt"

func main() {
	x1todo := newTodo("x1Go")
	fmt.Printf("Todo: %+v\n", x1todo)

	x1todo.updateTitle("x2Go")
	fmt.Printf("Todo: %+v\n", x1todo)

	x1todo.updateCompleted()
	fmt.Printf("Todo: %+v\n", x1todo)

	var todos Todos
	todos.appendTodo(&x1todo)
	fmt.Println(todos.format())

	todos.deleteTitle(0)
	fmt.Println(todos.format())
}
