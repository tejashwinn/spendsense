# Spendsense Backend Dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o spendsense ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/spendsense .
ENV PORT=8080
EXPOSE 8080
CMD ["./spendsense"]
