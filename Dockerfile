FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-gateway ./cmd/gateway/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/api-gateway .

COPY .env .

EXPOSE 8080

CMD ["./api-gateway"]
