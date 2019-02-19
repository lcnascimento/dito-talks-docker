FROM golang:1.11.4-stretch

WORKDIR /dito-talks-docker

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /dito-talks-docker

WORKDIR /dito-talks-docker/cmd/seeder
RUN ["go", "build"]

WORKDIR /dito-talks-docker/cmd/server
RUN ["go", "build"]

ENTRYPOINT [ "/dito-talks-docker/cmd/server/server" ]