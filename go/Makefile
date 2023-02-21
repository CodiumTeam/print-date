default: docker-test

# Local commands
.PHONY: test
test:
	go test -v ./...

.PHONY: docker-test
docker-test:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make test
