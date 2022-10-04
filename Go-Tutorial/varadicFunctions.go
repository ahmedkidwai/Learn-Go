package main

import "fmt"

func getSum2(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(getSum2(1, 2, 3, 4, 5))
}
