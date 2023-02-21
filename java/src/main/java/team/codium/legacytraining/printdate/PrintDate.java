package team.codium.legacytraining.printdate;

public class PrintDate {

	private final Calendar calendar;
	private final Printer printer;

	public PrintDate(Calendar calendar, Printer printer) {
		this.calendar = calendar;
		this.printer = printer;
	}

	public void printCurrentDate() {
		String line = calendar.today().toString();
		printer.printLine(line);
	}

}
