# Goal
Be able to test printCurrentDate function without changing the method signature.

1. Test the code with doubles using a library.
2. Test the code with doubles created by you.

# Code to test
    public function printCurrentDate()
    {
        $now = $this->calendar->now();
        $this->printer->printLine($now);
    }
    
## Run the kata
On Linux and Mac

    make

in Windows using docker

    docker run --rm -v ${PWD}:/code codiumteam/legacy-training-php:print-date make

In case of problems, please read the [Makefile](./Makefile)

# Learnings
How to use Prophecy to generate the doubles.

How to build a Mock and Stub manually.

## Tools
[Prophecy](https://github.com/phpspec/prophecy)
### Example of stub with prophecy

    /** @test */
    function should_success_when_password_is_valid()
    {
        $prophecy = $this->prophesize(PasswordValidator:class);
        $prophecy->isValid(‘aPassword’)->willReturn(true);
        /** @var PasswordValidator $passwordValidator */
        $passwordValidator = $prophecy->reveal();
        $userRegistration = new UserRegistration($passwordValidator);
    
        $success = $userRegistration->run();
    
        $this->assertTrue($success);
    }

	
### Example of spy with prophecy
    /** @test */
    function should_send_an_email()
    {
        $prophecy = $this->prophesize(EmailSender:class);
        /** @var EmailSender $emailSender */
        $emailSender = $prophecy->reveal();
        $userRegistration = new UserRegistration($emailSender);
    
    
        $userRegistration->register();
    
    
        $prophecy->send()->shouldHaveBeenCalled();
    }

## Author
Codium [@codiumteam](https://www.twitter.com/codiumteam) [www.codium.team](https://www.codium.team)