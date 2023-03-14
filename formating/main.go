package main

import "fmt"

type A struct {
	value int
	b     *B
}

type B struct {
	value string
}

func main() {
	tuple := A{
		value: int(^uint(0) >> 1),
		b: &B{
			value: "utc",
		},
	}

	fmt.Printf("%T: %+v\n", tuple, tuple)
}
