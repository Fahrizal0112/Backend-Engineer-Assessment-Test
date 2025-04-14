FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/banking-service ./cmd/api

FROM alpine:3.18

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/banking-service /app/banking-service
COPY .env /app/.env

EXPOSE 8080

ENTRYPOINT ["/app/banking-service"]
