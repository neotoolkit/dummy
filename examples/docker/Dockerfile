FROM golang:latest

RUN go install github.com/go-dummy/dummy/cmd/dummy@latest

COPY ./openapi3.yml .

WORKDIR .

CMD ["dummy", "-port=8080", "server", "openapi3.yml"]
