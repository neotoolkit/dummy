FROM golang:latest

RUN git clone https://github.com/go-dummy/dummy.git

COPY ./openapi3.yml ./dummy

WORKDIR ./dummy

RUN go mod download

CMD ["go", "run", "./cmd/dummy", "server", "openapi3.yml", "--port", "8080"]