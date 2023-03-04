# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Build the application in the cmd folder
RUN go build -o /go/bin/app ./cmd/

# Set the entrypoint to the app binary
ENTRYPOINT ["/go/bin/app"]