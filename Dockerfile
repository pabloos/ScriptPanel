FROM golang:stretch AS tester

RUN mkdir -p /go/src/ScriptPanel/scriptpanel
WORKDIR /go/src/ScriptPanel/scriptpanel

ENV GO111MODULE=on

COPY . .

RUN go mod download
####
FROM golang:stretch AS builder

RUN mkdir -p /go/src/ScriptPanel/scriptpanel
WORKDIR /go/src/ScriptPanel/scriptpanel

ENV GO111MODULE=on

COPY . .

# RUN go get github.com/docker/docker/client && \
#     go get github.com/gorilla/mux && \
#     go get github.com/jlaffaye/ftp && \
#     go get gopkg.in/ldap.v2 && \
#     go get gopkg.in/mgo.v2

RUN go mod download

CMD GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o arlequin-latest-linux scriptpanel/cmd/webserver/main.go