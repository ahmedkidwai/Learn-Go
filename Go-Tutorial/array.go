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
	fmt.Println("The length of the array is: ", len(arr2))

	// Let's loop an array with an iterator
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	// Let's loop an array with a range
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	fmt.Println("Multi Dimenonal Array")
	arr3 := [2][2]int{
		{1, 4},
		{3, 2},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr3[i][j])
		}
	}

	// Slice

	aStr1 := "abcd"
	rArr := []rune(aStr1)

	for _, v := range rArr {
		fmt.Println("Rune Array: ", v)
	}

	byteArr := []byte{'a', 'b', 'c', 'd'}
	bStr := string(byteArr[:])
	fmt.Println(bStr)

}
