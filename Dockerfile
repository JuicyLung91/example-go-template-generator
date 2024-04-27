FROM golang:latest

WORKDIR /go/src/app

RUN go mod init test && \
    go get -u -v github.com/gobeam/stringy

COPY . .

RUN go build -o main .

CMD ["./main"]
