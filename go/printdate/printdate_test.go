package printdate_test

import (
	"printdate/printdate"
	"testing"
)

func Test_PrintDate(t *testing.T) {
	printDate := printdate.NewPrintDate(
		printdate.NewCalendar(),
		printdate.NewPrinter(),
	)

	printDate.PrintCurrentDate()

	// How can we test this function?
}
