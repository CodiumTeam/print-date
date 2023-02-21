import Foundation

public struct PrintDate {
    private let calendar: Calendar;
    private let printer: Printer;

    public init(calendar: Calendar, printer: Printer) {
        self.calendar = calendar;
        self.printer = printer;
    }

    public func print_current_date() {
        let format = DateFormatter()
        format.timeStyle = .long
        format.dateStyle = .none
        let today = self.calendar.today()
        self.printer.print_line(line: format.string(from: today))
    }
}
