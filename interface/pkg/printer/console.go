package printer

import "fmt"

type ConsolePrinter struct {
}

func NewConsolePrinter() *ConsolePrinter {
	return &ConsolePrinter{}
}

func (io *ConsolePrinter) Println(a string) error {
	fmt.Println(a)
	return nil
}
