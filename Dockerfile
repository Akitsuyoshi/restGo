FROM golang:1.8.3

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/Akitsuyoshi/restGo/

# Install our dependencies
RUN go get github.com/gorilla/mux

# Install api binary globally within container
RUN go install github.com/Akitsuyoshi/restGo/

#

# Set binary as entrypoint
ENTRYPOINT /go/bin/restGo


# Expose default port (8080)
EXPOSE 8080
