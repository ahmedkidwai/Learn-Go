package main

import "fmt"

func main() {

	// Default value for array: 0, 0.0, false, true

	var arr1 [5]int
	arr1[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2[3])

	myVal := arr1[0]
	fmt.Println(myVal)
}
