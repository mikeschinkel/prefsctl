package errutil

import (
	"bytes"
	"fmt"
)

var _ Printer = (*BufferPrinter)(nil)

type BufferPrinter struct {
	bytes.Buffer
}

func (b *BufferPrinter) Print(i ...any) {
	b.WriteString(fmt.Sprint(i...))
}

func (b *BufferPrinter) Println(i ...any) {
	b.Print(i...)
	b.WriteByte('\n')
}

func (b *BufferPrinter) Printf(format string, i ...any) {
	b.WriteString(fmt.Sprintf(format, i...))
}
