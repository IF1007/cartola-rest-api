FROM golang:1.7
ENV GOPATH /go
COPY . /go/src/github.com/dijckstra/cartola-rest-api

RUN cd /go/src/github.com/dijckstra/cartola-rest-api && go get ./... && go build -o main.go

CMD ["/go/bin/cartola-rest-api"]