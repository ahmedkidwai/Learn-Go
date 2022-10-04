package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		// if nil then no error
		fmt.Println("Hi ", name)
	} else {
		log.Fatal(err)
	}
}
