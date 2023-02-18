FROM golang:1.20-alpine3.17

WORKDIR /build

COPY go.mod go.mod
COPY main.go main.go

RUN go build .

ENTRYPOINT [ "./tr-action" ]