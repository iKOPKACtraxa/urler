# Builder stage
FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /app

# Copy Go mod and sum files
# COPY go.mod go.sum ./
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app ./

# Set the entry point
ENTRYPOINT ["./app"]
