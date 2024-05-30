FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go --parseDependency --parseInternal
RUN go build -o main .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

CMD ["./main"]