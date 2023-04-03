package main

import (
	"fmt"
)

func main() {
	var attribute = fmt.Sprintf("output: %t", true)
	fmt.Println(attribute)
	fmt.Printf("%T\n", attribute)
	fmt.Printf("%T\n", true)
}
