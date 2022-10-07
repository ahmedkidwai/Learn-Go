package main

import "fmt"

func main() {

	// var myMap map [keyType]valueType

	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Spiderman"] = "Peter Parker"
	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krpyto", 2: "Bat Hound"}
	fmt.Println("Batman is: ", heroes["Batman"])
	// Will return nil
	fmt.Println("Chip is: ", superPets[3])
	_, ok := superPets[3]
	fmt.Println("Is there a third pet: ", ok)

	// Prints Key then value!!
	for k, v := range heroes {
		fmt.Println(k, v)
	}

	// delete
	delete(heroes, "Spiderman")
	for k, v := range heroes {
		fmt.Println(k, v)
	}

}
