FROM golang:1.15.6-alpine3.13

# Install Git
RUN apk update && apk add git

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/jimvmathews/generator

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENTRYPOINT [ "generator" ]