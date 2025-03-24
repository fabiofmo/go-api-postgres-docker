FROM golang:1.18

#Set working directory
#/go/src/app
#/app
WORKDIR /Users/fabio/OneDrive/Documents/DEV/Go/go-api-postgres-docker


# Copy the source code
COPY . .  

# EXPOSE the port on container
EXPOSE 8080

# Build the Go app
RUN go build -o main ./cmd/main.go

# Run the executable
CMD ["./main"]