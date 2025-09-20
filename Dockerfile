# Spendsense Backend Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o spendsense ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/spendsense .
COPY ./migrations ./migrations
ENV PORT=8080
EXPOSE 8080
CMD ["./spendsense"]
