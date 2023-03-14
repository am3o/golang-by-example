package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func main() {
	var w = new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 0, 0, ' ', 0)
	fmt.Fprintln(w, "------------------------------------------------------------------------")
	fmt.Fprintln(w, "Code\tLength\t\t\tPath")
	fmt.Fprintln(w, "------------------------------------------------------------------------")

	for i := 0; i < 3; i++ {
		fmt.Fprintln(w, fmt.Sprintf("%d\t%v\t\t%s\t", i, math.Pi, "/foo/bar"))
	}
	w.Flush()
}
