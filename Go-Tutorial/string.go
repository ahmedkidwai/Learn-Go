package main

import (
	"fmt"
	"strings"
)

func main() {

	// Just an array of bytes
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	fmt.Println(sV2)

	// Get length
	fmt.Println("Length", len(sV2))
	// Check if contained
	fmt.Println("Contains", strings.Contains(sV2, "Another"))
	fmt.Println("o index: ", strings.Index(sV2, "o"))

	sV3 := "\t Yes Hello World  HI \n"
	fmt.Printf("%q", sV3)
	sV4 := strings.TrimSpace(sV3)
	fmt.Printf("%q", sV4)
}
