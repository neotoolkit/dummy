FROM golang:latest

RUN go install github.com/go-dummy/dummy/cmd/dummy@latest

COPY ./openapi3.yml .

WORKDIR .

CMD ["dummy", "server", "openapi3.yml", "-port=8080"]
