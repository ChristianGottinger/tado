FROM golang:1.18-alpine

################## Compile  ##############################
ADD . /src
WORKDIR /src/cmd

RUN go build -o tado
RUN cp tado $GOPATH/bin

WORKDIR /src/cmd
