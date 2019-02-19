FROM node:9.5.0-alpine

WORKDIR /dito-talks-docker

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY ./src /dito-talks-docker/src

ENTRYPOINT [ "npm", "start" ]
