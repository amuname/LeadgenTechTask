FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.16

WORKDIR /root/

COPY --from=builder /app/docs ./docs

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
