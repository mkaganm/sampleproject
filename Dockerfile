# Build stage
FROM golang:1.20-alpine as builder
WORKDIR /app

# Copy necessary files and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code and build the application
COPY . .
RUN go build -o main .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the compiled application and .env file
COPY --from=builder /app/main .
COPY .env .

# Start the application using the environment variable for port
CMD ["sh", "-c", "./main"]
