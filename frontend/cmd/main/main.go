//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Coipo frontend started at", time.Now().Format(time.RFC3339))
	fmt.Println("Visit http://localhost:8080 to access the application.")
	// Here you would typically start your web server or application logic.
	// For example, you might use http.ListenAndServe to start a web server.
	// This is just a placeholder for demonstration purposes.
	select {} // Block forever
}
