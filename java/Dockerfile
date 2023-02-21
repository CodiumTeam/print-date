FROM gradle:7.4.1-jdk17-alpine

RUN apk add --no-cache make

ENV GRADLE_USER_HOME /gradle-cache
ENV GRADLE_OPTS "-Dorg.gradle.daemon=false"

WORKDIR /code

VOLUME ["/code"]
