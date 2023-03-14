package main

import (
	"flag"
	"fmt"
)

func main() {
	parameter := flag.String("opt", "", "unused flag parameter")
	flag.Parse()

	fmt.Println(parameter)
}
