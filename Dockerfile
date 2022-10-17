FROM golang:1.18-alpine3.16
RUN apk add build-base

RUN mkdir -p /go/src/github.com/wednesday-solutions/go-template-mysql

ADD . /go/src/github.com/wednesday-solutions/go-template-mysql

WORKDIR /go/src/github.com/wednesday-solutions/go-template-mysql

RUN GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go mod vendor
RUN mkdir -p /go/src/github.com/wednesday-solutions/go-template-mysql/output
RUN go build -o ./output/main ./cmd/server/main.go
RUN go build -o ./output/seeder ./cmd/seeder/main.go
RUN go build -o ./output/migrations ./cmd/migrations/main.go

FROM golang:1.18-alpine3.16
ARG ENVIRONMENT_NAME
RUN mkdir -p /app/scripts
RUN mkdir -p /app/internal/migrations
RUN mkdir -p /app/cmd/seeder
ADD .  /app

WORKDIR /app

COPY --from=0 /go/src/github.com/wednesday-solutions/go-template-mysql/output /app/
COPY --from=0 /go/src/github.com/wednesday-solutions/go-template-mysql/cmd/seeder /app/cmd/seeder
COPY --from=0 /go/src/github.com/wednesday-solutions/go-template-mysql/internal/migrations /app/internal/migrations
COPY --from=0 /go/src/github.com/wednesday-solutions/go-template-mysql/scripts /app/scripts
COPY --from=0 /go/src/github.com/wednesday-solutions/go-template-mysql/.env* /app/

CMD ["sh", "/app/scripts/migrate-and-run.sh"]
EXPOSE 9000