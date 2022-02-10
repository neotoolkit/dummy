FROM golang:latest

RUN go install github.com/neotoolkit/dummy/cmd/dummy@latest
