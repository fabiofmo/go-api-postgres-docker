FROM golang:1.24.1

#Set working directory
#/go/src/app
WORKDIR /go/src/app


# Copy the source code
COPY . .

# EXPOSE the port on container
EXPOSE 8080

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]