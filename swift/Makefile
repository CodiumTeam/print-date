# To check if Makefile has correct syntax (all must be tabs (^I), no spaces)
# cat -e -t -v Makefile

default: docker-test

.PHONY: test
test:
	swift test

.PHONY: coverage
coverage:
	swift test --enable-code-coverage
	llvm-cov show .build/*/debug/*.xctest -instr-profile .build/*/debug/codecov/default.profdata --format=html > coverage-report.html
	llvm-cov report .build/*/debug/*.xctest -instr-profile .build/*/debug/codecov/default.profdata -use-color
	@printf "\nOpen coverage-report.html\n"


.PHONY: docker-test
docker-test:
	docker run -it -v ${PWD}:/code -w /code swift:5.6 make test
