package main

import (
	"fmt"
	"testing"

	"github.com/am3o/golang-by-example/interface/pkg/printer"
	"github.com/stretchr/testify/assert"
)

func TestFooTypeConsole(t *testing.T) {
	s := NewFooType(printer.NewConsolePrinter(), "")

	assert.NoError(t, s.FooFunc())
}

func TestBarFunc(t *testing.T) {
	s := NewFooType(nil, "BarFunc")

	assert.Equal(t, s.BarFunc(), "BarFunc")
}

func TestFooFunc(t *testing.T) {
	value := "FooFunc"
	testPrinter := printer.NewTestPrinter()
	s := NewFooType(testPrinter, value)
	s.FooFunc()

	assert.Equal(t, testPrinter.Value(), fmt.Sprintf("%v hello from FooFunc", value))
}

func TestFooFuncTestTable(t *testing.T) {
	tt := []struct {
		value         string
		expectedValue string
	}{
		{"", " hello from FooFunc"},
		{"foo", "foo hello from FooFunc"},
		{"bar", "bar hello from FooFunc"},
	}

	for _, tc := range tt {
		t.Run(tc.value, func(t *testing.T) {
			testPrinter := printer.NewTestPrinter()
			s := NewFooType(testPrinter, tc.value)
			s.FooFunc()

			assert.Equal(t, testPrinter.Value(), tc.expectedValue)
		})
	}
}
