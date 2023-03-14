package main

import "fmt"

func Error() (err error) {
	defer func() {
		fmt.Println(err)
	}()
	return fmt.Errorf("could not start new session")
}

func main() {
	defer fmt.Println("another defer function")
	if err := Error(); err != nil {
		fmt.Printf("main session: %v\n", err)
	}
}
