package main

import "fmt"

func sayHello() {
	fmt.Println("Say Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {

	if y == 0 {
		return 0, fmt.Errorf("you can't divide by 0")
	} else {
		return x / y, nil
	}
}
func main() {
	// func funcName(param) return type { body }
	sayHello()
	fmt.Println(getSum(1, 2))
	fmt.Println(getTwo(3))

	fmt.Println("Division w/ Error Handling")
	fmt.Println("First dividing without an error")
	result, err := getQuotient(4, 5)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Println("Next dividing with an error (dividing by 0)")
	result2, err := getQuotient(4, 0)

	if err == nil {
		fmt.Println(result2)
	} else {
		fmt.Println(err)
	}

}
