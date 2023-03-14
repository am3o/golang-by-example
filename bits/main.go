package main

import (
	"fmt"
)

func main() {
	var x int64 = 0xAC
	fmt.Printf("%b\n", x)
	// Bit Masking
	x &= 0xF0
	fmt.Printf("%b\n", x)

	// bit shifting
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)

	// bit division
	c := 200
	fmt.Printf("%d\n", c>>1)

	// bit multiply
	fmt.Printf("%d\n", c<<1)

	// bit examples
	x = (x << 1) + 1
	fmt.Printf("%08b\n", x)
	x &= 255
	fmt.Printf("%08b\n", x)
	y := x & 64
	fmt.Printf("%08b\n", y)

	fmt.Printf("%b\n", ToByte([]int{2, 1, 23, 0}))
	fmt.Printf("%b\n", ToByte([]int{2, 1, 24, 0}))
	fmt.Printf("%b\n", ToByte([]int{1, 2, 24, 0}))
	fmt.Printf("%b\n", ToByte([]int{2, 2, 1, 1}))
	fmt.Printf("%b\n", ToByte([]int{3, 0, 0, 0}))
	fmt.Printf("%b\n", ToByte([]int{0, 0, 1, 0}))
	fmt.Printf("%b\n", ToByte([]int{0, 0, 2, 0}))
	fmt.Printf("%b\n", ToByte([]int{0, 1, 0, 1}))
}

func ToByte(version []int) (result int) {
	for _, o := range version {
		result = (result << 16) + o
	}
	return
}
