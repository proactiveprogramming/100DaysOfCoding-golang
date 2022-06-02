// File: helloworldinit/helloworldinit.go
// Version: 0.1
// -------------------------------------------------------
package main

import (
	"fmt"
	"localhost/myproject/greetingsinit"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
