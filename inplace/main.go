package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		fmt.Printf("\r%30s\r%d \r", " ", i)
	}
	os.Exit(-1)
}
