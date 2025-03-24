FROM golang:1.24.1

#Set working directory
#/go/src/app
#/Users/fabio/OneDrive/Documents/DEV/Go/go-api-postgres-docker
WORKDIR /app


# Copy the source code
COPY . .  

# EXPOSE the port on container
EXPOSE 8080

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]