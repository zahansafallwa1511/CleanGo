# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main ./cmd/main.go

# Production stage
FROM scratch

COPY --from=builder /app/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]
