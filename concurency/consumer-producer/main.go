package main

import (
	"fmt"
	"strconv"
	"sync"
)

func Producer(name string, n int, c chan string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		c <- fmt.Sprintf("name: %s produce: %v", name, i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	c := make(chan string, 1)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go Producer(strconv.Itoa(i), 4, c, &wg)
	}

	go func(pipe chan string) {
		for elem := range pipe {
			fmt.Printf("consume: %v\n", elem)
		}
	}(c)

	wg.Wait()
	close(c)
}
