class PrintDate {
  constructor(calendar, printer) {
    this.calendar = calendar;
    this.printer = printer;
  }

  printCurrentDate() {
    const today = this.calendar.today();
    this.printer.printLine(today);
  }
}

export default PrintDate;
