package main

import (
	"os"
)

func main() {
	if err := os.WriteFile("/tmp/greeting.txt", []byte("Hello, Gophers!"), 0747); err != nil {
		panic(err)
	}
}
