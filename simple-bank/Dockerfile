# Build stage

FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go
RUN apk add --no-cache make
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


# Run Stage

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
