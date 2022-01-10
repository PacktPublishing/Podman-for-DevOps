# Builder image 
FROM docker.io/library/golang 
 
# Copy files for build 
COPY go.mod /go/src/hello-world/ 
COPY main.go /go/src/hello-world/ 
 
# Set the working directory 
WORKDIR /go/src/hello-world 
 
# Download dependencies 
RUN go get -d -v ./... 
 
# Install the package 
RUN go build -v 
 
# Runtime image 
FROM registry.access.redhat.com/ubi8/ubi-micro:latest 
COPY --from=0 /go/src/hello-world/hello-world / 
EXPOSE 8080 
 
CMD ["/hello-world"]
