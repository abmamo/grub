# Base image
FROM golang:alpine
# Linux depdendencies
RUN apk update
RUN apk upgrade
RUN apk add bash
RUN apk add --update make
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Move to working directory /build
WORKDIR /build
# Copy files to working directory /build
COPY . /build
# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Install linter
RUN go get -u golang.org/x/lint/golint
# Run tests
RUN make test
# Run linting
RUN make vet
RUN make lint
# Build the application
RUN make build
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
# Copy environment file
COPY .env .
# Copy binary from build to main folder
RUN cp /build/grubAPI .
# Remove build files
RUN rm -r /build
# Export necessary port
EXPOSE 3333
# Command to run when starting the container
CMD ["/dist/grubAPI"]