package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	output, err := exec.Command("go", "env").Output()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(output))
}
