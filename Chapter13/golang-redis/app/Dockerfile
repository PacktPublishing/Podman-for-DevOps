FROM docker.io/library/golang AS builder

# Copy files for build
RUN mkdir -p /go/src/golang-redis
COPY go.mod main.go /go/src/golang-redis

# Set the working directory
WORKDIR /go/src/golang-redis

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go build -v 

# Runtime image
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest as bin
COPY --from=builder /go/src/golang-redis/golang-redis /usr/local/bin
COPY entrypoint.sh /

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
