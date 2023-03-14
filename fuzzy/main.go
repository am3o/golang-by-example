package main

import "fmt"

func Transform(value string) (string, error) {
	switch value {
	case "BaR":
		return "", fmt.Errorf("could not handle this input: %v", value)
	default:
		return value, nil
	}
}

func main() {
	fmt.Println(Transform("FooBar"))
}
