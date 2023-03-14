package main

import "fmt"

func main() {
	for i := 1; i <= 10; i -= -1 {
		defer fmt.Printf("%d\n", i)
		fmt.Printf("%d\n", i)
	}
	fmt.Println("The end?")
}
