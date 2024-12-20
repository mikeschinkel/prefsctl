package macprefs

import (
	"fmt"
)

var _ Printer = (*StandardPrinter)(nil)

type StandardPrinter struct{}

func (s StandardPrinter) Print(i ...any) {
	fmt.Print(i...)
}

func (s StandardPrinter) Println(i ...any) {
	fmt.Println(i...)
}

func (s StandardPrinter) Printf(format string, i ...any) {
	fmt.Printf(format, i...)
}
