package main

import "fmt"

func getArraySum(arr []int) int {

	sum := 0

	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {

	myIntArr := []int{1, 2, 3, 4, 5}
	// Passed by value. Any changes to the array in getArraySum will not affect the original array.
	fmt.Println(getArraySum(myIntArr))
}
