package main

import "fmt"

func main() {
	// var name []datatype
	// Usually will use make to make a slice.
	sl1 := make([]string, 6)
	sl1[0] = "Ahmed"
	sl1[1] = "is"
	sl1[2] = "the"
	sl1[3] = "very"
	sl1[4] = "best"
	sl1[5] = "person"
	fmt.Println("Slice Size: ", len(sl1))

	for _, v := range sl1 {
		fmt.Println(v)
	}

	// A slice is a view of an underlying array.

	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[:3]
	fmt.Println(sl3)

	sArr[0] = 10
	// sl3 has a pointer to sArr to sArr so we expect it to update too.
	fmt.Println(sl3)

	sl3 = append(sl3, 12)
	fmt.Println(sl3)
	fmt.Println(sArr)

	// creating an empty string slice has a value of nil
	sl4 := make([]string, 6)
	fmt.Println(sl4)

}
