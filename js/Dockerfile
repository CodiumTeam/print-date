FROM node:17-alpine

MAINTAINER Codium <info@codium.team>

RUN apk update && apk add make && rm -rf /var/cache/apk/*

WORKDIR /code

ADD ./package*.json ./
RUN npm install

USER node
RUN npm config set update-notifier false

VOLUME ["/code", "/code/node_modules"]

