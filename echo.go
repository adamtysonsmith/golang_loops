package main

import (
  "fmt"
  "os"
)

func main() {
  // Declare two vars of type string
  var printString, seperator string
  
  // Loop!
  // The shorthand i := 0 can only be used inside a function, and must be used as the initializer in a for loop (can't use var keyword)
  // The command line args are accessed using os.Args
  // len(os.Args) gives us the length of command line arguments
  // we start at 1 because the first arg is a path..
     // os.Args[0] - name of the command that it is running as
     // os.Args[1] - first command line parameter, ...
  var length int = len(os.Args)
  
  for i := 1; i < length; i++ {
    printString += seperator + os.Args[i]
    seperator = " "
  }
  
  fmt.Println(printString)
  
  
  // For loops are versatile, because they don't require each part (initializer, condition, increment)
  // You could simply do this to run an infinite loop
  /*
    for {
         some inifine code
    } 
  */
  
  // A classic while loop
  var a = 1;
  
  for a < 10 {
    fmt.Println("Around the loop..")
    a += 1
  }
  
  
  // Range Loops
  var printMe string
  var seper8 string
  
  // We use _ as a blank identifier if we don't need to access the key in the range
  // Go does not allow unused variables like temp, that's why we need it
  // this range iterates over a slice, starting at 1
  for _, arg := range os.Args[1:] {
    printMe += seper8 + arg
    seper8 = " "
  }
  
  fmt.Println(printMe)
  
}