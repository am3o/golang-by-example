package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzTransform(f *testing.F) {
	tt := []struct {
		input          string
		expectedResult string
	}{
		{
			input:          "",
			expectedResult: "",
		},
		{
			input:          "Foo",
			expectedResult: "Foo",
		},
		{
			input:          "Bar",
			expectedResult: "Bar",
		},
		{
			input:          "Buzz",
			expectedResult: "Buzz",
		},
	}

	for _, tc := range tt {
		f.Add(tc.input)
	}

	f.Fuzz(func(t *testing.T, input string) {
		actual, err := Transform(input)

		assert.NoError(t, err)
		assert.Equal(t, input, actual)
	})
}
