#####################################
#   STEP 1 build executable binary  #
#####################################
FROM golang:1.16.2-alpine AS builder

# create appuser
RUN adduser -D -g '' elf

# create workspace
WORKDIR /opt/app/
COPY go.mod .
COPY go.sum .

# fetch dependancies
RUN go mod download
RUN go mod verify

# copy the source code as the last step
COPY . .

# build binary
RUN go build -o /go/bin/ ./pkg/...
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/ ./pkg/...

#####################################
#   STEP 2 build a small image      #
#####################################
FROM alpine:3.12.3

# arg variable to identify the service name
ARG SERVICE_NAME

# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd

# copy the static executable
COPY --from=builder --chown=elf:1000 /go/bin/${SERVICE_NAME} /main

# use a non-root user
USER elf

# run app
ENTRYPOINT ["./main"]