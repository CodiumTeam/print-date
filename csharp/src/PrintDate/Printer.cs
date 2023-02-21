namespace PrintDate;

public class Printer : IPrinter
{
    public void PrintLine(string line)
    {
        Console.WriteLine(line);
    }
}