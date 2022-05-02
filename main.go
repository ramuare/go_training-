/*package main

import "fmt"

func main() {
    var age int // variable declaration
    fmt.Println("My age is", age)
}*/

package main

import "go_training/go_training-/Cal"

func main() {
	Cal.Add(23, 55)
	Cal.Sub(44, 65)
	Cal.Mult(5, 4)
	Cal.Div(10, 2)
}
