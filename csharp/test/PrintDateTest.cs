using NSubstitute;
using PrintDate;

namespace KataPrintDate.Tests
{
    public class PrintDateTest
    {
        [Fact]
        public void think_a_good_name_and_change_it()
        {
            var printDate = new PrintDate.PrintDate(new Calendar(), new Printer());

            printDate.PrintCurrentDate();

            // I don't know how to test it 
        }
    }
}
