FROM golang:1.23-alpine

RUN apk --no-cache add ca-certificates gcc g++ libc-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service ./cmd/main

CMD ["/app/service"]
