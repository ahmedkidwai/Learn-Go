package main

import (
	"fmt"
)

func main() {
	fmt.Println("For Loops")

	// Notice semi colon here
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	// While loop uses for syntax
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// Range
	aNums := []int{1, 2, 3}
	// _ is saying we don't care about index. We could add an index though
	for _, num := range aNums {
		fmt.Println(num)
	}

}
