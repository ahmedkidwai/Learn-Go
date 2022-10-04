package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	cV1 := 1.5
	cV2 := int(cV1)
	// Will round down
	fmt.Println(cV2)

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3)

	// 5000000, nil, intg
	fmt.Println(cV4, err, reflect.TypeOf(cV4))
}
