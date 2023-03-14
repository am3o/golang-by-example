package main

import (
	"fmt"

	"github.com/elgs/gojq"
)

var jsonObj = `
{
  "name": "sam",
  "gender": "m",
  "pet": null,
  "skills": [
    "Eating",
    "Sleeping",
    "Crawling"
  ],
  "hello.world":true
}
`

func main() {
	parser, err := gojq.NewStringQuery(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Query("name"))          // sam <nil>
	fmt.Println(parser.Query("gender"))        // m <nil>
	fmt.Println(parser.Query("skills.[2]"))    // Sleeping <nil>
	fmt.Println(parser.Query("hello"))         // <nil> hello does not exist.
	fmt.Println(parser.Query("pet"))           // <nil> <nil>
	fmt.Println(parser.Query("."))             // map[name:sam gender:m pet:<nil> skills:[Eating Sleeping Crawling] hello.world:true] <nil>
	fmt.Println(parser.Query("'hello.world'")) // true <nil>
}
