package main

import "fmt"

func main() {

	// Condiitonal operators: > > >= <= == !=
	// Logical Operators: && ||

	iAge := 8

	if (iAge >= 0) && (iAge <= 18) {
		fmt.Println("A baby")
	} else {
		fmt.Println("Not a baby")
	}
}
