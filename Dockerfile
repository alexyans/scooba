FROM golang:1.10.1 AS builder

RUN apt-get update && apt-get install -y libgit2-dev

RUN go get github.com/urfave/cli
RUN go get -v -x gopkg.in/libgit2/git2go.v24

ADD . /go/src/github.com/alexyans/scooba/
WORKDIR /go/src/github.com/alexyans/scooba/
