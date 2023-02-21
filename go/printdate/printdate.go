package printdate

type PrintDate interface {
	PrintCurrentDate()
}

func NewPrintDate(calendar Calendar, printer Printer) PrintDate {
	return printDate{calendar: calendar, printer: printer}
}

type printDate struct {
	calendar Calendar
	printer  Printer
}

func (printDate printDate) PrintCurrentDate() {
	today := printDate.calendar.today()
	printDate.printer.printLine(today.String())
}
