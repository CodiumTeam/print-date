package printdate

import "fmt"

type Printer interface {
	printLine(line string)
}

func NewPrinter() Printer {
	return printer{}
}

type printer struct {
}

func (printer printer) printLine(line string) {
	fmt.Println(line)
}
