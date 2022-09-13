FROM ubuntu:22.04

RUN apt update && \
    apt install golang-go -y

COPY foo.go .
RUN go build ./foo.go

CMD ./foo