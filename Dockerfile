# Stage 1 — Build the app
FROM golang:1.23.3-alpine AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the binary
RUN go build -o main .

# Stage 2 — Run the app
FROM alpine:latest

WORKDIR /app

# Copy only the built binary from stage 1
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 3000

CMD ["./main"]