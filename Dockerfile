FROM golang:1.19-alpine

WORKDIR /app

ENV GIN_MODE release

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY internal ./internal
COPY Makefile ./

RUN go get github.com/mattn/go-isatty@v0.0.16
RUN go build -o /go-template-api

EXPOSE 8081

CMD [ "/go-template-api" ]