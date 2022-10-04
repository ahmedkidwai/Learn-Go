package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rStr := "abcdefg"
	fmt.Println("Rune Count: ", utf8.RuneCountInString(rStr))

	// i is index and runeVal is individual
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)

	}

	// Expected output
	// Rune Count:  7
	// 0 : U+0061 'a' : a
	// 1 : U+0062 'b' : b
	// 2 : U+0063 'c' : c
	// 3 : U+0064 'd' : d
	// 4 : U+0065 'e' : e
	// 5 : U+0066 'f' : f
	// 6 : U+0067 'g' : g

}
