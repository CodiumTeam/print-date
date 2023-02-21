# Goal
Be able to test printCurrentDate function without changing the method signature.

1. Test the code with doubles using a library.
2. Test the code with doubles created by you.

# Code to test
    func (printDate PrintDate) printCurrentDate() {
        today := printDate.calendar.today()
        printDate.printer.printLine(today.String())
    }
## Run the kata
On Linux and Mac

    make

in Windows using docker

    docker run --rm -v ${PWD}:/code -w /code golang:1.18 make test

In case of problems, please read the [Makefile](./Makefile)

# Learnings

How to use TBD to generate the doubles.

How to build a Mock and Stub manually.

## Author
Codium [@codiumteam](https://www.twitter.com/codiumteam) [www.codium.team](https://www.codium.team)