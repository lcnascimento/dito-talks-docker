FROM golang:1.11.4-stretch as builder

WORKDIR /dito-talks-docker

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . /dito-talks-docker

ENV CGO_ENABLED=0

WORKDIR /dito-talks-docker/cmd/seeder
RUN ["go", "build"]

WORKDIR /dito-talks-docker/cmd/server
RUN ["go", "build"]

FROM scratch

COPY --from=builder /dito-talks-docker/cmd/server/server /go/bin/server
COPY --from=builder /dito-talks-docker/cmd/seeder/seeder /go/bin/seeder

CMD [ "/go/bin/server" ]