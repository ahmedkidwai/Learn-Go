package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a basic overview of variables")

	// var name type
	var vName string = "Ahmed"
	fmt.Println(vName)

	var v1, v2 = 1.2, 3.4
	// Asking compiler to figure out  the type
	var v3 = "Hello"

	// Cleanest with type detection
	v4 := 2.4
	// If you were to re-assign v4 later you'd use
	v4 = 3.0

}
