# Goal
Unit test the printCurrentDate function.

1. Test the code with doubles from a library.
2. Test the code with doubles created by you.

# Code to test


```swift
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
```
## Run the kata
On Linux and Mac

make

in Windows using docker

docker run -it -v ${PWD}:/code -w /code swift:5.6 make test

In case of problems, please read the [Makefile](./Makefile)

# Learnings
How to build a Mock and Stub manually.

How to use Unittest Mock to generate the doubles.

## Tools

    [mockingbird](https://github.com/birdrides/mockingbird)

### Example of spy


```swift
    // Mocking
    let bird = mock(Bird.self)

    // Verification
    verify(bird.fly()).wasCalled()
```


### Example of stub


```swift
    // Mocking
    let bird = mock(Bird.self)

    // Stubbing
    given(bird.canFly).willReturn(true)

```

## Author
Codium [@codiumteam](https://www.twitter.com/codiumteam) [www.codium.team](https://www.codium.team)