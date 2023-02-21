namespace PrintDate;

public class Calendar : ICalendar
{
    public DateTime Today()
    {
        return DateTime.Now;
    }
}