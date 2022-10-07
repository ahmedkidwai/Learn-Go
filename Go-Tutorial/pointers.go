package main

import "fmt"

func changeValue(ptr *int) {
	*ptr = 12

}

func main() {

	f := 10
	fmt.Println("f before func: ", f)
	changeValue(&f)
	fmt.Println("f after func: ", f)

	var f4Ptr *int = &f
	fmt.Println("The address of f4", f4Ptr)
	fmt.Println("The value of f4", *f4Ptr)
}
