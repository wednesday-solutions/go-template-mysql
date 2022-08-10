#!/bin/sh
set -a && source .env.local && set +a

export MYSQL_HOST=localhost
# drop first and then run migrations

go run ./cmd/migrations/main.go down

# seed data
go run ./cmd/seeder/main.go
