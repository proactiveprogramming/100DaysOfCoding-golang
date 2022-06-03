// File: hello/hello.go
// Version: 1.0
// Original author: Melissa Neman (Proactive Programming)
// Purpose: Minimal amount for a legal golang main application that runs on the command line.
// -------------------------------------------------------
package main

import (
  "fmt"
  "rsc.io/quote"
)

func main() {
  fmt.Println("Hello World!")

  fmt.Println(quote.Go())
}
