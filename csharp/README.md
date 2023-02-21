# Goal
Be able to test printCurrentDate function without changing the method signature.

1. Test the code with doubles using a library.
2. Test the code with doubles created by you.

# Code to test
    public void PrintCurrentDate()
    {
        var line = _calendar.Today().ToString();
        _printer.PrintLine(line);
    }

## Run the kata
On Linux and Mac

    make

in Windows using docker

    docker run --rm -v ${PWD}:/code codiumteam/legacy-training-csharp:print-date make test

In case of problems, please read the [Makefile](./Makefile)
### Example of Mock/Spy

	[Fact]
	public void should_interact_with_the_mock() {
    	collaborator = Substitute.For<Collaborator>();
    	MyClass myClass = new MyClass(collaborator);

    	myClass.Run();

    	collaborator.Received().Collaborate();
	}
### Example of Stub

	[Fact]
	public void should_retrieve_the_stub_response(){
    	collaborator = Substitute.For<Collaborator>();
    	String response = "collaborator response";
    	calculator.Collaborate().Returns(response);
    	MyClass myClass = new MyClass(collaborator);

    	String result = myClass.Run();

    	Assert.Equal(response, result);
	}

## Author
Codium [@codiumteam](https://www.twitter.com/codiumteam) [www.codium.team](https://www.codium.team)
