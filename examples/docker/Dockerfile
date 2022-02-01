FROM golang:latest

RUN go install github.com/go-dummy/dummy/cmd/dummy@latest

COPY ./openapi3.yml .

WORKDIR .

CMD ["dummy", "s", "openapi.yml", "-port=8080"]
