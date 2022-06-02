// File: greetingsmulti/greetingsmulti.go
// Version: 0.1
// -------------------------------------------------------
// Purpose: To create a greetings package.
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
