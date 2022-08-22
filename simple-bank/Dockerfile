# Install tools stage
FROM alpine AS tools
WORKDIR /tools
RUN apk add --no-cache curl
# Need to install this way because the run image does not have go installed, so different strategy used than migrate dockerfile. Also I am not sure where go installs it exactly. It might work, likely just in usr/bin. This is easier for now
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
# Build stage

FROM golang:1.17-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go


# Run Stage

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=tools /tools/migrate /usr/bin/migrate
COPY .env ./.env
COPY start.sh .
COPY db/migrations ./db/migrations

RUN chmod u+x ./start.sh

EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
