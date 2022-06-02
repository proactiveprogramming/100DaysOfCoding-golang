// File: helloworldmulti/helloworldmulti.go
// Version: 1.0
// -------------------------------------------------------
package main

import (
	"fmt"
	"localhost/myproject/greetingsmulti"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
