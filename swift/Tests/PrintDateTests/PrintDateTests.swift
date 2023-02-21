import XCTest
@testable import PrintDate

final class PrintDateTests: XCTestCase {
    func test_xxx() throws {
        let print_date = PrintDate(calendar: Calendar(), printer: Printer());
        
        print_date.print_current_date()

        // I don't know how to test it
    }
}
