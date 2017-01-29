FROM golang:1.7.4

RUN set -x && \
    apt-get update && \
    apt-get install -y netcat && \
    apt-get clean

COPY . /go/watcher

ENV GOPATH=/go/watcher

WORKDIR /go/watcher

EXPOSE 8080:8080

# RUN go build main.go && ./main
