FROM golang:latest

RUN git clone https://github.com/go-dummy/dummy.git

COPY ./openapi3.yml ./dummy

WORKDIR ./dummy

RUN go mod download

CMD ["go", "run", "./cmd/dummy", "-port=8080", "server", "openapi3.yml"]
