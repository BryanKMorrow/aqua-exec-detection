############################
# STEP 1 build executable binary
############################
FROM golang:rc-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Copy code

RUN mkdir -p $GOPATH/src/github.com/BryanKMorrow/aqua-exec-detection
ADD . $GOPATH/src/github.com/BryanKMorrow/aqua-exec-detection
WORKDIR $GOPATH/src/github.com/BryanKMorrow/aqua-exec-detection

# Fetch dependencies.
# Using go get.
RUN go get "github.com/gorilla/mux"; go get "github.com/gorilla/handlers"; go get "github.com/parnurzeal/gorequest";

# Build the binary.
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o $GOPATH/src/github.com/BryanKMorrow/aqua-exec-detection/aqua-exec-detection cmd/aqua-exec-detection/main.go

############################
# STEP 2 build a small image
############################
FROM scratch
# Import from the builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable.
COPY --from=builder /go/src/github.com/BryanKMorrow/aqua-exec-detection/aqua-exec-detection /go/src/github.com/BryanKMorrow/aqua-exec-detection/aqua-exec-detection

# Run the binary.
ENTRYPOINT ["/go/src/github.com/BryanKMorrow/aqua-exec-detection/aqua-exec-detection"]
