import PrintDate from '../src/print_date';
import Calendar from '../src/calendar';
import Printer from '../src/printer';

describe('PrintDate', function () {
  it('change_this_name', function () {
    let printDate = new PrintDate(new Calendar(), new Printer());

    printDate.printCurrentDate();

    // I don't know how to test it
  });
});
