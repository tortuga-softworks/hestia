FROM golang:1.20-alpine

WORKDIR /hestia-src

# Downloading dependencies
COPY . .
RUN go mod download

# Building the application
RUN go build -o /hestia ./cmd

# Using a new base image to run the binary
FROM alpine:latest  

WORKDIR /root/

COPY --from=0 /hestia /hestia

CMD [ "/hestia" ]
