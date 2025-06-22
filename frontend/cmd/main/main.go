//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"time"
)

func main() {
	go func ()  {
		for{
			time.Sleep(1 * time.Second)
			fmt.Println("Hello from WebAssembly in a goroutine!")
		}	
	}()
	fmt.Println("Hello from WebAssembly!")

	select {} // Block forever to keep the program running
}
