FROM golang:1.20.0-bullseye as build

RUN apt-get update && \
    apt-get install lsb-release -y

RUN go version

RUN apt-get update

ADD . /app
WORKDIR /app/cmd

RUN go build -o /out/tado

WORKDIR /app
RUN go vet ./...
RUN go test ./... --cover

#############################################################
FROM gcr.io/distroless/base:debug AS final

USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /out /out

WORKDIR /out

ENTRYPOINT ["./tado"]