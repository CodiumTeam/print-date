package team.codium.legacytraining.printdate;

import org.junit.jupiter.api.Test;
import team.codium.legacytraining.printdate.Calendar;
import team.codium.legacytraining.printdate.PrintDate;
import team.codium.legacytraining.printdate.Printer;


public class PrintDateTest {
	@Test
	public void printDate() throws Exception {
		PrintDate printDate = new PrintDate(new Calendar(), new Printer());

		printDate.printCurrentDate();

		// How can we test this function?
	}
}
