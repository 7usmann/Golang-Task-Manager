# Step 1: Use the official Go image as a base
FROM golang:1.19-alpine AS builder

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy the application files
COPY . .

# Step 6: Build the application
RUN go build -o app main.go

# Step 7: Use a smaller image for the final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY templates /app/templates
COPY static /app/static
COPY migrations /app/migrations

# Expose the port on which the application runs
EXPOSE 8080

# Command to run the app
CMD ["./app"]
