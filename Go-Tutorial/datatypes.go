package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int float64, bool, string, rune
	// Default: 0, 0,0, false, "", ""

	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(3.31))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf('🎉'))
}
