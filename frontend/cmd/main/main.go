package main

import (
	"fmt"
	"shared/todo"
	"syscall/js"
)

func main() {

	c := make(chan struct{}, 0)

	// Get the document object
	document := js.Global().Get("document")

	// Create a new paragraph element
	p := document.Call("createElement", "p")

	mytodo := todo.Todo{
		ID:        "1",
		Title:     "Learn Go",
		Completed: false,
	}

	p.Set("textContent", fmt.Sprintf("Todo: %s, Completed: %t", mytodo.Title, mytodo.Completed))

	// Append the paragraph to the body
	document.Get("body").Call("appendChild", p)

	<-c
}
