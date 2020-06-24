FROM golang:latest

RUN mkdir /web

ADD . /web/

WORKDIR /web

RUN go build -o main .

CMD ["/web/main"]
