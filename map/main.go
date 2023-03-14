package main

import "fmt"

func main() {
	c := map[string]string{
		"bar": "foo",
	}
	c["foo"] = "bar"
	delete(c, "foo")
	fmt.Printf("%+v", c)

}
