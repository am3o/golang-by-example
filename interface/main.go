package main

import "github.com/am3o/golang-by-example/interface/pkg/printer"

type Printer interface {
	Println(a string) error
}

type FooType struct {
	printer Printer
	value   string
}

func NewFooType(p Printer, fooBar string) FooType {
	return FooType{
		printer: p,
		value:   fooBar,
	}
}

func (t FooType) FooFunc() error {
	return t.printer.Println(t.value + " hello from FooFunc")
}

func (t FooType) BarFunc() string {
	return "BarFunc"
}

func main() {
	s := NewFooType(printer.NewConsolePrinter(), "BarFunc")
	s.FooFunc()
	s.BarFunc()
}
