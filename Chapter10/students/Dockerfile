FROM docker.io/library/golang AS builder

# Copy files for build
RUN mkdir -p /go/src/students/models
COPY go.mod main.go /go/src/students
COPY models/main.go /go/src/students/models

# Set the working directory
WORKDIR /go/src/students

# Download dependencies
RUN go get -d -v ./...

# Install the package
RUN go build -v 

# Runtime image
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest as bin
COPY --from=builder /go/src/students /usr/local/bin
COPY entrypoint.sh /

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
