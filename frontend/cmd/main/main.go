package main

import (
	"fmt"
	"shared/todo"
)

func main() {
	mytodo := todo.Todo{
		ID:        "1",
		Title:     "Learn Go",
		Completed: false,
	}

	fmt.Println("Todo ID:", mytodo.ID)
}
