# Development Dockerfile
FROM golang:1.24

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN sqlc generate
EXPOSE 8080
CMD ["go", "run", "cmd/app/main.go"]
