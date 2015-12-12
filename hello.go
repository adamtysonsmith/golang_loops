// The main package is special, it defines a standalone executable program
package main

// We also import whatever other packages we want using 'import'
// You must import all dependencies exactly, you cannot import extra packages that arent used!
import (
  "fmt"
)

// The main function is also special, where the execution of the program beings
func main() {
  fmt.Println("Hello, Boulder!")
}