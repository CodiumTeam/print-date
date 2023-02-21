.PHONY: default test docker-test docker-build
default: docker-test

test:
	./gradlew :test

docker-build:
	docker build -f ../Dockerfile -t codiumteam/legacy-training-java .

docker-test:
	docker run --rm -v ${PWD}:/opt/project -v ${PWD}/.gradle:/gradle-cache -w /opt/project codiumteam/legacy-training-java make test
