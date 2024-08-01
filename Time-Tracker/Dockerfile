# Use the official Golang image to create a build artifact
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download


COPY . .

RUN go build -o main .


FROM alpine:latest


WORKDIR /root/


COPY --from=builder /app/main .


EXPOSE 8080


CMD ["./main"]
