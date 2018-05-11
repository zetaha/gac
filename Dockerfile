FROM golang:latest

WORKDIR /go/src/app
ADD ./ /go/src/app

RUN go get -d -v github.com/gorilla/mux
RUN go build -o main

ENTRYPOINT ["/go/src/app/main"]
