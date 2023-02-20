default: docker-tests

# Local commands
dependencies:
	composer install
.PHONY: tests
tests:
	./vendor/bin/phpunit
coverage:
	./vendor/bin/phpunit --coverage-text

# Docker commands
docker-build:
	docker build -t codiumteam/legacy-training-php:print-date -f Dockerfile .

docker-tests:
	@docker run --rm -v ${PWD}:/code codiumteam/legacy-training-php:print-date make tests
docker-coverage:
	@docker run --rm -v ${PWD}:/code codiumteam/legacy-training-php:print-date make coverage

