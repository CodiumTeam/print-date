namespace PrintDate;

public class PrintDate
{
    private readonly ICalendar _calendar;
    private readonly IPrinter _printer;

    public PrintDate(ICalendar calendar, IPrinter printer)
    {
        _calendar = calendar;
        _printer = printer;
    }

    public void PrintCurrentDate()
    {
        var line = _calendar.Today().ToString();
        _printer.PrintLine(line);
    }
}
