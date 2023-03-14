package main

import (
	"fmt"
	"testing"
)

const serialized = `{
  "version": 42
}`

var result string

func BenchmarkHash(b *testing.B) {
	for index, input := range []string{
		"",
		"foo",
	} {
		for _, bc := range []func(string) string{
			hashMD5,
			hashHighway64,
			hashHighway256,
		} {
			b.Run(fmt.Sprintf("%v", index%3), func(b *testing.B) {
				var r string
				for n := 0; n < b.N; n++ {
					r = bc(input)
				}
				result = r
			})
		}
	}

}
