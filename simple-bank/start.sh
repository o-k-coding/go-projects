#!/bin/sh

set -e

echo "load env variables"
source /app/.env

echo "run db migration"
migrate -path /app/db/migrations -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB?sslmode=$POSTGRES_SSL_MODE" -verbose up

echo "start the app"
exec "$@"
