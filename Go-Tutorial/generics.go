package main

import "fmt"

// NOTE GENERICS ARE NEW. YOU NEED GO 1.18 or > to use them.
type MyContraint interface {
	int | float64
}

// The generic basically says we expect our variables to be of T type.
func getSum23[T MyContraint](x T, y T) T {

	return x + y
}

func main() {

	fmt.Println("5+4 = ", getSum23(5, 4))
	fmt.Println("5.5+4.3= ", getSum23(5.5, 4.3))
	// This won't work because you can't mix.
	//fmt.Println("5+4.3= ", getSum23(5, 4.3))
	// This won't work because you can't use a string.
	//./generics.go:21:37: default type rune of '5' does not match inferred type int for T
	//fmt.Println("5+'5'= ", getSum23(5, '5'))
}
