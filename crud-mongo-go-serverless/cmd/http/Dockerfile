#TODO : Custom it and runn only targeted docker file
# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application for a Linux environment
# The output is named 'bootstrap' as expected by the AWS Lambda custom runtime
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /bootstrap ./cmd/http/main.go

# Stage 2: Create the final runtime image
FROM public.ecr.aws/lambda/provided:al2023

# Copy the compiled binary from the builder stage
COPY --from=builder /bootstrap /var/task/bootstrap

# Ensure the bootstrap file is executable
RUN chmod 755 /var/task/bootstrap

# Set the command to run when the container starts
ENTRYPOINT [ "/var/task/bootstrap" ]
