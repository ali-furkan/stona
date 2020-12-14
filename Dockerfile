FROM golang:latest

WORKDIR /app

ARG ENV_PATH=.env

COPY ${ENV_PATH} .

COPY go.mod go.sum Makefile ./

RUN make install

COPY main.go .

COPY auth auth
COPY config config
COPY images images
COPY storages storages
COPY tools tools

ENV PORT=8080

RUN make build

EXPOSE ${PORT}

CMD ["./main"]