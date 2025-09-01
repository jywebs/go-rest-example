FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM busybox:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
