FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/hafizbadrie/journal-app
COPY . .

RUN go build -o /go/bin/journal-app
EXPOSE 8080
CMD ["/go/bin/journal-app"]
