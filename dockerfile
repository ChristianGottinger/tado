FROM golang:1.19-alpine

################## Compile  ##############################
ADD . /src
WORKDIR /src/cmd

RUN go build -o tado
RUN cp tado $GOPATH/bin

WORKDIR /src/cmd
