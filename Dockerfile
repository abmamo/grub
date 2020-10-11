FROM golang:alpine

RUN apk update
RUN apk upgrade
RUN apk add bash

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o grubAPI .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
COPY .env /dist

# Copy binary from build to main folder
RUN cp /build/grubAPI .

# Export necessary port
EXPOSE 3333

# Command to run when starting the container
CMD ["/dist/grubAPI"]