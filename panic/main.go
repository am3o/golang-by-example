package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		time.Sleep(time.Second / 2)
		panic("Test panic!")
	}()

	time.Sleep(time.Second)
	fmt.Println("done!")
}
