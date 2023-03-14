package main

import "fmt"

func main() {
	x := []rune("ho")
	fmt.Printf("%s\n", string(x))

	fmt.Println(Insert(string(x), 1, "ell"))
}

func Insert(original string, pos int, insert string) string {
	if len(original) < pos {
		pos = len(original)
	}

	return original[:pos] + insert + original[pos:]
}
