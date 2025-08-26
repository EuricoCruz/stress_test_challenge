FROM golang:1.25.0-alpine3.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o stress_test_challenge ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stress_test_challenge .

RUN apk add --no-cache ca-certificates
ENTRYPOINT ["./stress_test_challenge"]

docker run stress_test —url=http://google.com —requests=100 —concurrency=10
