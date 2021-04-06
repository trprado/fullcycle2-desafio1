FROM golang:1.16-alpine

WORKDIR /go/src


COPY src/* .

ENV GOOS=linux

ENTRYPOINT ["go", "run", "http-server.go"]
EXPOSE 8000
