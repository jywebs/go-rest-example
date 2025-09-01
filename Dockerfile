FROM golang:1.20-alpine AS builder

WORKDIR /app
# Install required tools
RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Copy source code
COPY . .

# Generate swagger documentation
RUN swag init --parseDependency --parseInternal -g cmd/server/main.go

# Download dependencies, verify and build
RUN go mod tidy && \
    go mod verify && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM busybox:latest

WORKDIR /app
# Copy both the server binary and the docs directory
COPY --from=builder /app/server .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
CMD ["./server"]
