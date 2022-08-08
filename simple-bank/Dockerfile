# Build stage

FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go

# Run Stage

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY .env.docker ./.env

EXPOSE 8080

CMD [ "/app/main" ]
