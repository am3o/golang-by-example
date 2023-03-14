package main

import (
	"errors"
	"fmt"
)

var (
	innerError = errors.New("some error happened")
)

func NoOp() error {
	return fmt.Errorf("%w: connection error: %w", fmt.Errorf("%w", innerError), fmt.Errorf("inner error"))
}

func main() {
	err := NoOp()
	if err != nil {
		switch {
		case errors.Is(err, innerError):
			fmt.Printf("first case: %v", err)
		default:
			fmt.Printf("default: %v", err)
		}
	}
}
