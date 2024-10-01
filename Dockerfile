# Use Go image
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the rest of the source code
COPY . .

# Initialize the module and fetch dependencies
RUN go mod init yask-tracker && go mod tidy

# Compile the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o yask-tracker cmd/tracker/main.go

# Use a smaller image to run the binary
FROM alpine:latest

# Copy the binary from the build stage
COPY --from=builder /app/yask-tracker /usr/local/bin/yask-tracker

# Set the entry point
ENTRYPOINT ["yask-tracker"]
