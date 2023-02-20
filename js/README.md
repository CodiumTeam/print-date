# Print date
## Goal
Be able to test printCurrentDate function without changing the method signature.

1. Test the code with doubles using a library.
2. Test the code with doubles created by you.

## Code to test
	public void printCurrentDate() {
		String line = calendar.today().toString();
		printer.printLine(line);
	}

## Learnings
* How to use Sinon to generate the doubles.
* How to build a Mock and Stub manually.

### Tools
[Sinon](https://sinonjs.org/)

#### Example of Spy	

    it(‘should send an email’, function()
    {
        let emailSender = new EmailSender();
        let spy = sinon.spy(emailSender, ‘send’);
        let userRegistration = new UserRegistration(emailSender);

        userRegistration.run();

        assert(spy.calledOnce);
    }
#### Example of Stub    
    it(‘should success when password is valid’, function()
    {
        let passwordValidator = new PasswordValidator();
        sinon.stub(passwordValidator, ‘isValid’).returns(true);
        let userRegistration = new UserRegistration(passwordValidator);
    
        let success = userRegistration.run();
    
    
        assert(success);
    }

## How to run the tests
### With Docker
Run `make`

### Without Docker
1. `npm install`
2. `npm test`

## Author
Codium [@codiumteam](https://www.twitter.com/codiumteam) [www.codium.team](https://www.codium.team)